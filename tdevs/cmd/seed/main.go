package main

import (
	"fmt"
	"tdevs/server/profile"
	"tdevs/store/db"
)

func main() {
	profile, _ := profile.GetProfile()
	dbDriver, _ := db.NewDBDriver(profile)

	fmt.Println(dbDriver)

}
