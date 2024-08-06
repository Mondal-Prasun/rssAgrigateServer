-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    publishedAt TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feedId UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;