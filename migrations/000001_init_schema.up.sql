CREATE TABLE "product" (
  "id" bigserial PRIMARY KEY,
  "manufacturer_id" bigint NOT NULL, 
  "name" varchar NOT NULL,
  "price" decimal(12, 2) NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "attribute_values" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "attribute_id" bigint NOT NULL,  
  "value" varchar NOT NULL  
);

CREATE TABLE "attributes" (
  "id" bigserial PRIMARY KEY,
  "category_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "type" smallint NOT NULL    
);

CREATE TABLE "product_images" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "image" varchar NOT NULL,
  "title" varchar NOT NULL  
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "parent_id" int NOT NULL DEFAULT '0',
  "category_name" varchar NOT NULL,
  "category_description" varchar NOT NULL  
);

CREATE TABLE "product_category" (  
  "product_id" bigint NOT NULL,
  "category_id" bigint NOT NULL   
);

CREATE TABLE "manufacturers" (
  "id" bigserial PRIMARY KEY,  
  "name" varchar NOT NULL,
  "description" varchar NOT NULL  
);


ALTER TABLE "product" ADD FOREIGN KEY ("manufacturer_id") REFERENCES "manufacturers" ("id");

ALTER TABLE "product_properties" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
ALTER TABLE "product_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

CREATE INDEX ON "product" ("name");

CREATE INDEX ON "categories" ("category_name");

CREATE INDEX ON "manufacturers" ("name");