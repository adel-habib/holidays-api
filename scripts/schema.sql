CREATE TABLE "regions" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255),
  "short_name" varchar(255),
  "parent_id" bigint
);

ALTER TABLE "regions" ADD FOREIGN KEY ("parent_id") REFERENCES "regions" ("id");