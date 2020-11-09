CREATE TABLE IF NOT EXISTS "product" (
  "id" bigserial PRIMARY KEY,
  "manufacturer_id" bigint NOT NULL, 
  "category_id" bigint NOT NULL, 
  "name" varchar(100) NOT NULL,
  "slug" varchar(20) NOT NULL,
  "price" decimal(12, 2) NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "product" ADD FOREIGN KEY ("manufacturer_id") REFERENCES "manufacturers" ("id");
ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
CREATE INDEX ON "product" ("name");

CREATE TABLE IF NOT EXISTS "attribute_values" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "attribute_id" bigint NOT NULL,  
  "value_int" int NULL,  
  "value_str" varchar(100) NULL  
);

ALTER TABLE "attribute_values" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
ALTER TABLE "attribute_values" ADD FOREIGN KEY ("attribute_id") REFERENCES "attributes" ("id");

CREATE TABLE IF NOT EXISTS "attributes" (
  "id" bigserial PRIMARY KEY,
  "category_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "slug" varchar(20) NOT NULL,
  "type" smallint NOT NULL    
);

CREATE TABLE IF NOT EXISTS "product_images" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "url" varchar(100) NOT NULL,
  "title" varchar(100) NOT NULL  
);

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

CREATE TABLE IF NOT EXISTS "categories" (
  "id" bigserial PRIMARY KEY,  
  "category_name" varchar(100) NOT NULL,
  "slug" varchar(20) NOT NULL,
  "category_description" varchar(100) NOT NULL  
);

CREATE INDEX ON "categories" ("category_name");

CREATE TABLE IF NOT EXISTS "manufacturers" (
  "id" bigserial PRIMARY KEY,  
  "name" varchar(100) NOT NULL,
  "description" varchar(100) NOT NULL  
);

CREATE INDEX ON "manufacturers" ("name");