CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "nickname" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "country" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);