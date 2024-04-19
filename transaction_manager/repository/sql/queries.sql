-------------------------------------------------
-- TRANSACTION TABLE
-------------------------------------------------

-- name: GetTransactionById :one
SELECT *
FROM "Transaction"
WHERE "id" = $1;

-- name: GetTransactionsBySku :many
SELECT *
FROM "Transaction"
WHERE "sku" = $1;

-- name: GetTransactionsByAuthorId :many
SELECT *
FROM "Transaction"
WHERE "author_id" = $1;

-- name: GetTransactionsByAuthorName :many
SELECT *
FROM "Transaction"
WHERE "author_id" = (SELECT "id" FROM "User" WHERE "name" = $1);

-- name: GetTransactionsByStartAndEndDate :many
SELECT *
FROM "Transaction"
WHERE ("date" >= @start_date OR @start_date IS NULL)
  AND ("date" <= @end_date OR @end_date IS NULL);

-- name: CreateTransaction :one
INSERT INTO "Transaction" ("action", "sku", "shelf_name", "quantity", "author_id")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-------------------------------------------------
-- USER TABLE
-------------------------------------------------

-- name: GetUserById :one
SELECT *
FROM "User"
WHERE "id" = $1;

-- name: GetUserByExactName :one
SELECT *
FROM "User"
WHERE "name" = $1;

-- name: FindUserByName :many
SELECT *
FROM "User"
WHERE "name" ILIKE '%' || $1 || '%';

-- name: CreateUser :one
INSERT INTO "User" ("name")
VALUES ($1)
RETURNING *;
