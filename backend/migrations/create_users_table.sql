-- +migrate Up
CREATE TABLE "users"
(
    "id"        UUID PRIMARY KEY,
    "firstName" VARCHAR(50)  NOT NULL,
    "lastName"  VARCHAR(50)  NOT NULL,
    "createdAt" DATE         NOT NULL,
    "email"     VARCHAR(255) NOT NULL UNIQUE,
    "phone"     VARCHAR(16),
    "language"  VARCHAR(4),
    "disabled"  BOOLEAN      NOT NULL,
    "initPass"  BOOLEAN      NOT NULL
);

-- +migrate Down
DROP TABLE "users";