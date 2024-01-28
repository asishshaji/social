package postgresql

import (
	"context"
	"fmt"
	"log"
	"tdevs/store"
)

func (postgres *PostgresDB) GetGroups(ctx context.Context) []store.Group {
	groups := []store.Group{}

	err := postgres.db.Select(&groups, "SELECT * FROM groups")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return groups
}

func (postgres *PostgresDB) JoinGroup(ctx context.Context, g_id, u_id int32) error {
	_, err := postgres.db.Exec("INSERT INTO memberships(user_id, group_id) SELECT $1,$2 WHERE NOT EXISTS(SELECT 1 FROM memberships WHERE user_id = $3 AND group_id = $4)", u_id, g_id, u_id, g_id)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
