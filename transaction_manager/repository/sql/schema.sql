CREATE TABLE "User"
(
    "id"   integer PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "Transaction"
(
    "id"         integer PRIMARY KEY,
    "action"     integer   NOT NULL,
    "date"       timestamp NOT NULL DEFAULT now(),
    "sku"        varchar   NOT NULL,
    "shelf_name" varchar   NOT NULL,
    "quantity"   integer   NOT NULL CHECK ("quantity" > 0),
    "author_id"  integer   NOT NULL REFERENCES "User" ("id") ON DELETE CASCADE
);

COMMENT
    ON COLUMN "Transaction"."action" IS 'Action is stored as integer. Mapping between action and integer must be documented somewhere else.';

