CREATE TABLE IF NOT EXISTS "account"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "phone" VARCHAR(13) NOT NULL,
    "balance" NUMERIC,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);