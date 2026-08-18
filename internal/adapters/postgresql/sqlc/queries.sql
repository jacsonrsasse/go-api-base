-- name: ListSomething :many
SELECT *
  FROM something;

-- name: FindSomethingById :one
SELECT *
  FROM something
 WHERE id = $1;