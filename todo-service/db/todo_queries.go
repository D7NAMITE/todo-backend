package db

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type TodoRequest struct {
	ClerkID     string  `json:"UserID"`
	Title       string  `json:"Title"`
	Description *string `json:"Description"`
	Status      string  `json:"Status"`
}

func GetTodoByClerkID(clerkId string, databaseUrl string) ([]Todo, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return []Todo{}, err
	}
	queries := New(conn)

	return queries.GetTodoByClerkID(ctx, clerkId)
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

	err = queries.CreateTodoByClerkID(ctx, CreateTodoByClerkIDParams{
		ClerkID:     request.ClerkID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	})

	if err != nil {
		return err
	}

	return nil
}
