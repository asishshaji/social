package postgresql

import (
	"context"
	"fmt"
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
	_, err := postgres.db.Exec("INSERT INTO user_group_relation(user_id, group_id) VALUES($1,$2)", u_id, g_id)
	if err != nil {
		return err
	}

	return err
}
