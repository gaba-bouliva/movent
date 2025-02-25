-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: GetMovies :many
SELECT * FROM movies
ORDER BY title;

-- name: CreateMovie :one
INSERT INTO movies (
  title, year, runtime, genres
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateMovie :exec
UPDATE movies
  set title = $2,
  year = $3,
  runtime = $4,
  genres = $5
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM movies
WHERE id = $1;