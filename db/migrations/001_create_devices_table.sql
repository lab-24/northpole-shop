-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "device_types" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_time" int DEFAULT (0)
);

CREATE TABLE "locations" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_time" bigint DEFAULT (0)
);

CREATE TABLE "devices" (
  "id" uuid PRIMARY KEY,
  "device_type_id" uuid NOT NULL,
  "location_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "serial_number" varchar NOT NULL,
  "created_time" bigint DEFAULT (0)
);


CREATE INDEX ON "devices" ("device_type_id");

CREATE INDEX ON "devices" ("location_id");

ALTER TABLE "devices" ADD FOREIGN KEY ("device_type_id") REFERENCES "device_types" ("id");

ALTER TABLE "devices" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "devices";
DROP TABLE IF EXISTS "device_types";
DROP TABLE IF EXISTS "locations";
