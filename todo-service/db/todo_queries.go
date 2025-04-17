package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetTodoByUserId(userid string, databaseUrl string) ([]Todo, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return []Todo{}, err
	}
	queries := New(conn)

	uuid := pgtype.UUID{}
	err = uuid.Scan(userid)
	if err != nil {
		return []Todo{}, err
	}
	return queries.GetTodoByUid(ctx, uuid)
}
