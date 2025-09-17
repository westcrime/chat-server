-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE chats_users (
    id BIGSERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL
);

CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    text VARCHAR(511) NOT NULL,
    from_user VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_messages_chat_id ON messages(chat_id);
CREATE INDEX idx_messages_timestamp ON messages(timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS chats_users;
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
