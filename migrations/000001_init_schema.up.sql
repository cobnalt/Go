CREATE TABLE IF NOT EXISTS Product (
  id bigserial PRIMARY KEY,
  manufacturer_id bigint NOT NULL, 
  category_id bigint NOT NULL, 
  name varchar(100) NOT NULL,
  slug varchar(20) NOT NULL,
  price decimal(12, 2) NOT NULL,
  description text NOT NULL
);

CREATE TABLE IF NOT EXISTS Attribute_values (
  id bigserial PRIMARY KEY,
  product_id bigint NOT NULL,
  attribute_id bigint NOT NULL,  
  value_int int NULL,  
  value_str varchar(100) NULL  
);

CREATE TABLE IF NOT EXISTS Attributes (
  id bigserial PRIMARY KEY,
  category_id bigint NOT NULL,
  name varchar(100) NOT NULL,
  slug varchar(20) NOT NULL,
  attr_type smallint NOT NULL     
);

CREATE TABLE IF NOT EXISTS Product_images (
  id bigserial PRIMARY KEY,
  product_id bigint NOT NULL,
  url varchar(100) NOT NULL,
  title varchar(100) NOT NULL  
);

CREATE TABLE IF NOT EXISTS Categories (
  id bigserial PRIMARY KEY,  
  category_name varchar(100) NOT NULL,
  slug varchar(20) NOT NULL,
  category_description varchar(100) NOT NULL  
);

CREATE TABLE IF NOT EXISTS Manufacturers (
  id bigserial PRIMARY KEY,  
  name varchar(100) NOT NULL,
  description varchar(100) NOT NULL  
);

ALTER TABLE Product ADD FOREIGN KEY (manufacturer_id) REFERENCES Manufacturers (id);
ALTER TABLE Product ADD FOREIGN KEY (category_id) REFERENCES Categories (id);
ALTER TABLE Attribute_values ADD FOREIGN KEY (product_id) REFERENCES Product (id);
ALTER TABLE Attribute_values ADD FOREIGN KEY (attribute_id) REFERENCES Attributes (id);
ALTER TABLE Product_images ADD FOREIGN KEY (product_id) REFERENCES Product (id);
CREATE INDEX ON Product (name);
CREATE INDEX ON Manufacturers (name);
CREATE INDEX ON Categories (category_name);