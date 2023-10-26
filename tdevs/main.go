package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tdevs/server"
	"tdevs/server/profile"
	"tdevs/store"
	"tdevs/store/db"
)

func main() {
	profile, _ := profile.GetProfile()
	ctx, cancel := context.WithCancel(context.Background())
	dbDriver, _ := db.NewDBDriver(profile)
	store := store.NewStore(dbDriver, profile)
	s, err := server.NewServer(ctx, profile, store)
	if err != nil {
		cancel()
		fmt.Printf("error starting server%+v\n", err)
		return
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.Shutdown(ctx)
		cancel()
	}()

	if err := s.Start(ctx); err != nil {
		if err != http.ErrServerClosed {
			fmt.Printf("failed to start server: %v", err)
			cancel()
		}
	}

}
