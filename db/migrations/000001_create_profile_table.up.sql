CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "profile" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v1(),
    "first_name" VARCHAR(45) NOT NULL,
    "last_name" VARCHAR(45) NOT NULL,
    "email" VARCHAR(30) NOT NULL,
    "phone_number" VARCHAR(15) NOT NULL,
    "citizenship" VARCHAR(15) NOT NULL,
    "birth_date" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "birth_country" VARCHAR(30) NOT NULL,
    "residence_country" VARCHAR(30) NOT NULL,
    "password" VARCHAR (80) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "archived_at" TIMESTAMPTZ,
    CONSTRAINT "id_profile" PRIMARY KEY ("id")
)