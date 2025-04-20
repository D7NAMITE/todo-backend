-- +goose Up
CREATE TABLE "users" (
                         user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         clerk_id TEXT NOT NULL UNIQUE,
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "todos" (
                         todo_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         user_id UUID REFERENCES "users" (user_id),
                         title VARCHAR NOT NULL,
                         description VARCHAR,
                         status VARCHAR NOT NULL DEFAULT 'Todo',
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;
