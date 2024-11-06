-- Create "budgets" table
CREATE TABLE "public"."budgets" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "amount" numeric NULL DEFAULT 0,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_budgets_deleted_at" to table: "budgets"
CREATE INDEX "idx_budgets_deleted_at" ON "public"."budgets" ("deleted_at");
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_categories_deleted_at" to table: "categories"
CREATE INDEX "idx_categories_deleted_at" ON "public"."categories" ("deleted_at");
-- Create "reports" table
CREATE TABLE "public"."reports" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "report_data" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_reports_deleted_at" to table: "reports"
CREATE INDEX "idx_reports_deleted_at" ON "public"."reports" ("deleted_at");
-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NULL,
  "category_id" bigint NULL,
  "amount" numeric NULL,
  "transaction_date" timestamptz NULL,
  "type" text NULL,
  "description" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_transactions_deleted_at" to table: "transactions"
CREATE INDEX "idx_transactions_deleted_at" ON "public"."transactions" ("deleted_at");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "username" text NULL,
  "email" text NULL,
  "password" bytea NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_customer_email" to table: "users"
CREATE UNIQUE INDEX "idx_customer_email" ON "public"."users" ("email");
-- Create index "idx_username" to table: "users"
CREATE UNIQUE INDEX "idx_username" ON "public"."users" ("username");
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
