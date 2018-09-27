-- +migrate Up

CREATE TABLE "license_employee" (
  "license"    UUID      NOT NULL,
  "employee"   UUID      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  UNIQUE ("license", "employee")
);

CREATE TABLE "license_workplace" (
  "license"    UUID      NOT NULL,
  "workplace"  UUID      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  UNIQUE ("license", "workplace")
);



-- +migrate Down

DROP TABLE "license_workplace";

DROP TABLE "license_employee";
