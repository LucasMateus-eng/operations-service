BEGIN;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vehicle_licensing_status') THEN
        CREATE TYPE "vehicle_licensing_status" AS ENUM (
          'REGULAR',
          'LATE',
          'BLOCKED',
          'SEIZED',
          'STOLEN'
        );
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE "user_role" AS ENUM (
          'ADMINISTRATOR',
          'EMPLOYEE',
          'DRIVER'
        );
    END IF;
END
$$;

CREATE TABLE "drivers" (
  "id" int PRIMARY KEY,
  "name" text NOT NULL,
  "rg" text UNIQUE NOT NULL,
  "cpf" text UNIQUE NOT NULL,
  "driver_license" text UNIQUE NOT NULL,
  "date_of_birth" date NOT NULL,
  "cell_phone" text NOT NULL,
  "email" text NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "user_id" int UNIQUE NOT NULL
);

CREATE TABLE "vehicles" (
  "id" int PRIMARY KEY,
  "brand" text NOT NULL,
  "model" text NOT NULL,
  "year_of_manufacture" date NOT NULL,
  "plate" text UNIQUE NOT NULL,
  "renavam" text UNIQUE NOT NULL,
  "licensing_expiry_date" timestampz NOT NULL,
  "licensing_status" vehicle_licensing_status NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "drivers_vehicles" (
  "driver_id" int NOT NULL,
  "vehicle_id" int NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  PRIMARY KEY ("driver_id", "vehicle_id")
);

CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "username" text UNIQUE NOT NULL,
  "hashed_password" text NOT NULL,
  "role" user_role NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "adresses" (
  "id" int PRIMARY KEY,
  "locality" text NOT NULL,
  "number" text NOT NULL,
  "complement" text,
  "neighborhood" text NOT NULL,
  "city" text NOT NULL,
  "state" text NOT NULL,
  "cep" text NOT NULL,
  "country" text NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "user_id" int UNIQUE NOT NULL
);

CREATE INDEX "date_of_birth_index" ON "drivers" ("date_of_birth");

CREATE INDEX "rg_index" ON "drivers" ("rg");

CREATE INDEX "cpf_index" ON "drivers" ("cpf");

CREATE INDEX "driver_license_index" ON "drivers" ("driver_license");

CREATE INDEX "created_at_index" ON "drivers" ("created_at");

CREATE INDEX "updated_at_index" ON "drivers" ("updated_at");

CREATE INDEX "brand_index" ON "vehicles" ("brand");

CREATE INDEX "model_index" ON "vehicles" ("model");

CREATE INDEX "year_of_manufacture_index" ON "vehicles" ("year_of_manufacture");

CREATE INDEX "plate_index" ON "vehicles" ("plate");

CREATE INDEX "renavam_index" ON "vehicles" ("renavam");

CREATE INDEX "licensing_expiry_date_index" ON "vehicles" ("licensing_expiry_date");

CREATE INDEX "licensing_status_index" ON "vehicles" ("licensing_status");

CREATE INDEX "created_at_index" ON "vehicles" ("created_at");

CREATE INDEX "updated_at_index" ON "vehicles" ("updated_at");

CREATE INDEX "created_at_index" ON "drivers_vehicles" ("created_at");

CREATE INDEX "updated_at_index" ON "drivers_vehicles" ("updated_at");

CREATE INDEX "role_index" ON "users" ("role");

CREATE INDEX "created_at_index" ON "users" ("created_at");

CREATE INDEX "updated_at_index" ON "users" ("updated_at");

CREATE INDEX "city_index" ON "adresses" ("city");

CREATE INDEX "state_index" ON "adresses" ("state");

CREATE INDEX "cep_index" ON "adresses" ("cep");

CREATE INDEX "country_index" ON "adresses" ("country");

CREATE INDEX "created_at_index" ON "adresses" ("created_at");

CREATE INDEX "updated_at_index" ON "adresses" ("updated_at");

ALTER TABLE "drivers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "drivers_vehicles" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id") ON DELETE CASCADE;

ALTER TABLE "drivers_vehicles" ADD FOREIGN KEY ("vehicle_id") REFERENCES "vehicles" ("id") ON DELETE CASCADE;

ALTER TABLE "adresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

COMMIT;