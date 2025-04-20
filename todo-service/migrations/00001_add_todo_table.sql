-- +goose Up
CREATE TABLE "users" (
                       user_id uuid DEFAULT gen_random_uuid (),
                       username VARCHAR NOT NULL UNIQUE,
                       password VARCHAR NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW(),
                       PRIMARY KEY (user_id)
);

CREATE TABLE "todos" (
                       todo_id uuid DEFAULT gen_random_uuid (),
                       user_id uuid REFERENCES "users" (user_id),
                       title VARCHAR NOT NULL,
                       description VARCHAR NULL,
                       status VARCHAR NOT NULL DEFAULT 'Todo',
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW(),
                       PRIMARY KEY (todo_id)
);

-- +goose Down
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;
