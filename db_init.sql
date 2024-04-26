DROP TABLE IF EXISTS "Transaction";
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "Shelf";
DROP TABLE IF EXISTS "Product";

CREATE TABLE "Product"
(
    "id"       SERIAL PRIMARY KEY,
    "sku"      varchar UNIQUE NOT NULL,
    "name"     varchar        NOT NULL,
    "category" varchar
);

CREATE TABLE "Shelf"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       varchar NOT NULL,
    "product_id" integer NOT NULL REFERENCES "Product" ("id") ON DELETE CASCADE,
    "quantity"   integer NOT NULL CHECK ("quantity" > 0),

    UNIQUE ("name", "product_id")
);

CREATE UNIQUE INDEX ON "Product" ("sku");

CREATE TABLE "User"
(
    "id"   SERIAL PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "Transaction"
(
    "id"         SERIAL PRIMARY KEY,
    "action"     integer   NOT NULL,
    "date"       timestamp NOT NULL DEFAULT now(),
    "sku"        varchar   NOT NULL,
    "shelf_name" varchar   NOT NULL,
    "quantity"   integer   NOT NULL CHECK ("quantity" > 0),
    "author_id"  integer   NOT NULL REFERENCES "User" ("id") ON DELETE CASCADE
);

COMMENT
    ON COLUMN "Transaction"."action" IS 'Action is stored as integer. Mapping between action and integer must be documented somewhere else.';

INSERT INTO "User" ("name")
VALUES ('User1');
INSERT INTO "User" ("name")
VALUES ('User2');
INSERT INTO "User" ("name")
VALUES ('User3');
INSERT INTO "User" ("name")
VALUES ('User4');
INSERT INTO "User" ("name")
VALUES ('User5');

INSERT INTO "Product" ("sku", "name", "category")
VALUES ('sku1', 'Product1', 'category1');
INSERT INTO "Product" ("sku", "name", "category")
VALUES ('sku2', 'Product2', 'category2');
INSERT INTO "Product" ("sku", "name", "category")
VALUES ('sku3', 'Product3', 'category3');

INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf1', 1, 10);
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf2', 1, 10);
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf3', 1, 10);
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf2', 2, 10);
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf3', 2, 10);
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ('Shelf3', 3, 10);

INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku1', 'Shelf1', 10, 1);
INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku1', 'Shelf2', 10, 1);
INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku1', 'Shelf3', 10, 1);
INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku2', 'Shelf2', 10, 1);
INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku2', 'Shelf3', 10, 1);
INSERT INTO "Transaction" ("action", "date", "sku", "shelf_name", "quantity", "author_id")
VALUES (0, '2021-01-01 00:00:00', 'sku3', 'Shelf3', 10, 1);
