CREATE TABLE "Product"
(
    "id"           integer PRIMARY KEY,
    "sku"          varchar UNIQUE NOT NULL,
    "name"         varchar        NOT NULL,
    "expired_date" timestamp,
    "category"     varchar
);

CREATE TABLE "Shelf"
(
    "id"         integer PRIMARY KEY,
    "name"       varchar NOT NULL,
    "product_id" integer NOT NULL REFERENCES "Product" ("id") ON DELETE CASCADE,
    "quantity"   integer NOT NULL CHECK ("quantity" > 0),

    UNIQUE ("name", "product_id")
);

CREATE UNIQUE INDEX ON "Product" ("sku");

