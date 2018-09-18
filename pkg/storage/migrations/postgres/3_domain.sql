-- +migrate Up

CREATE TABLE "license_user" (
  "license"    UUID      NOT NULL,
  "user_id"    UUID      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  UNIQUE ("license", "user_id")
);

-- +migrate Down

DROP TABLE "license_user";
