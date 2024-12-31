-- Modify "budgets" table
ALTER TABLE "public"."budgets" ALTER COLUMN "user_id" TYPE text, ALTER COLUMN "category_id" TYPE text, ALTER COLUMN "amount" SET NOT NULL;
-- Modify "categories" table
ALTER TABLE "public"."categories" ALTER COLUMN "user_id" TYPE text;
-- Modify "reports" table
ALTER TABLE "public"."reports" ALTER COLUMN "user_id" TYPE text, ALTER COLUMN "category_id" TYPE text;
-- Modify "transactions" table
ALTER TABLE "public"."transactions" ALTER COLUMN "user_id" TYPE text, ALTER COLUMN "category_id" TYPE text;
