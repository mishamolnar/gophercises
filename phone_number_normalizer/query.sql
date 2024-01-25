-- name: GetNumber :one
SELECT * FROM NUMBERS
WHERE number = $1 LIMIT 1;

-- name: ListNumbers :many
SELECT * FROM NUMBERS
ORDER BY number
limit 10;

-- name: CreateNumber :one
INSERT INTO NUMBERS (
    input, number
) VALUES (
             $1, $2
         )
RETURNING input, number;

-- name: DeleteAuthor :exec
DELETE FROM NUMBERS
WHERE number = $1;