-- +goose Up
CREATE TABLE "todos" (
                         todo_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         clerk_id TEXT NOT NULL,
                         title VARCHAR NOT NULL,
                         description VARCHAR,
                         status VARCHAR NOT NULL DEFAULT 'Todo',
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose Down
DROP TABLE IF EXISTS todos;