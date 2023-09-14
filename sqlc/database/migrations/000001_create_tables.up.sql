CREATE TABLE IF NOT EXISTS "accounts" (
  "id" serial PRIMARY KEY,
  "owner" text NOT NULL,
  "balance" integer NOT NULL,
  "currency" text,
  "created_at" timestamptz DEFAULT NOW() NOT NULL 
);