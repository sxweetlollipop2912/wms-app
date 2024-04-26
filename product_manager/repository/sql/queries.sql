-------------------------------------------------
-- PRODUCT TABLE
-------------------------------------------------

-- name: GetProductById :one
SELECT *
FROM "Product"
WHERE "id" = $1
LIMIT 1;

-- name: GetProductBySku :one
SELECT *
FROM "Product"
WHERE "sku" = $1
LIMIT 1;

-- name: FindProductByName :many
SELECT *
FROM "Product"
WHERE "name" ILIKE '%' || $1 || '%';

-- name: GetProductByCategory :many
SELECT *
FROM "Product"
WHERE "category" = $1;

-- name: CreateProduct :one
INSERT INTO "Product" ("sku", "name", "category")
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :one
UPDATE "Product"
SET "sku"      = $1,
    "name"     = $2,
    "category" = $3
WHERE "id" = $4
RETURNING *;

-------------------------------------------------
-- SHELF TABLE
-------------------------------------------------

-- name: GetShelfById :one
SELECT *
FROM "Shelf"
WHERE "id" = $1
LIMIT 1;

-- name: GetShelfByProductId :many
SELECT *
FROM "Shelf"
WHERE "product_id" = $1;

-- name: GetShelfBySkuAndName :one
SELECT *
FROM "Shelf" AS "s"
WHERE s.product_id = (SELECT "id" FROM "Product" WHERE "sku" = $1 LIMIT 1)
  AND s.name = $2
LIMIT 1;

-- name: GetShelfBySku :many
SELECT *
FROM "Shelf"
WHERE "product_id" = (SELECT "id" FROM "Product" WHERE "sku" = $1 LIMIT 1);

-- name: CreateShelf :one
INSERT INTO "Shelf" ("name", "product_id", "quantity")
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateShelf :one
UPDATE "Shelf"
SET "name"       = $1,
    "product_id" = $2,
    "quantity"   = $3
WHERE "id" = $4
RETURNING *;

-- name: DeleteShelfById :one
DELETE
FROM "Shelf"
WHERE "id" = $1
RETURNING *;
