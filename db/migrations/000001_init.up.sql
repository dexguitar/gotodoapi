CREATE TABLE IF NOT EXISTS todos (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    done BOOLEAN NOT NULL
);

CREATE INDEX idx_todos_title ON todos(title);