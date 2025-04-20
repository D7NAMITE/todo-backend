package db

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type TodoRequest struct {
	UserID      string  `json:"UserID"`
	Title       string  `json:"Title"`
	Description *string `json:"Description"`
	Status      string  `json:"Status"`
}

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

func CreateTodo(r *http.Request, databaseUrl string) error {
	ctx := context.Background()

	var request TodoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return err
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
		}
	}(conn, ctx)

	queries := New(conn)

	uuid := pgtype.UUID{}
	err = uuid.Scan(request.UserID)
	if err != nil {
		return err
	}

	err = queries.CreateTodo(ctx, CreateTodoParams{
		UserID:      uuid,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	})

	if err != nil {
		return err
	}

	return nil
}
