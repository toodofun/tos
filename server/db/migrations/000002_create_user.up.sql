CREATE TABLE IF NOT EXISTS users
(
    id            TEXT         NOT NULL PRIMARY KEY,
    username      VARCHAR(32)  NOT NULL,
    password_hash VARCHAR(256) NOT NULL,
    email         VARCHAR(128) NOT NULL,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);