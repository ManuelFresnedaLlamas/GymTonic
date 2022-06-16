-- +migrate Up
CREATE TABLE "auth"
(
    "id"                 UUID PRIMARY KEY,
    "login"              VARCHAR(255) UNIQUE NOT NULL,
    "passwordHash"       VARCHAR(128)        NOT NULL,
    "passwordSalt"       VARCHAR(64)         NOT NULL,
    "passwordResetToken" VARCHAR(128),
    "role"               SMALLINT            NOT NULL,
    "userID"             UUID                NOT NULL,

    CONSTRAINT fk_user2auth FOREIGN KEY ("userID") REFERENCES "users" ("id") ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE "auth";