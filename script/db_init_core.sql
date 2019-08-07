CREATE TABLE IF NOT EXISTS users
(
    id        BIGSERIAL    NOT NULL
        CONSTRAINT users_pk PRIMARY KEY,
    email     VARCHAR(255) NOT NULL,
    username  VARCHAR(32)  NOT NULL,
    pass_hash VARCHAR(64)  NOT NULL,
    avatar    VARCHAR(255) DEFAULT ''
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_uindex
    ON users (LOWER(email));

CREATE UNIQUE INDEX IF NOT EXISTS users_username_uindex
    ON users (LOWER(username));

CREATE INDEX IF NOT EXISTS users_email_pass_hash_index
    ON users (LOWER(email), LOWER(pass_hash));

CREATE INDEX IF NOT EXISTS users_username_pass_hash_index
    ON users (LOWER(username), LOWER(pass_hash));
