// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: movies.sql

package data

import (
	"context"

	"github.com/lib/pq"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  title, year, runtime, genres
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, created_at, title, year, runtime, genres
`

type CreateMovieParams struct {
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime Runtime  `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
		arg.Title,
		arg.Year,
		arg.Runtime,
		pq.Array(arg.Genres),
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Title,
		&i.Year,
		&i.Runtime,
		pq.Array(&i.Genres),
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, id)
	return err
}

const getMovie = `-- name: GetMovie :one
SELECT id, created_at, title, year, runtime, genres FROM movies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Title,
		&i.Year,
		&i.Runtime,
		pq.Array(&i.Genres),
	)
	return i, err
}

const getMovies = `-- name: GetMovies :many
SELECT id, created_at, title, year, runtime, genres FROM movies
ORDER BY title
`

func (q *Queries) GetMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Title,
			&i.Year,
			&i.Runtime,
			pq.Array(&i.Genres),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :exec
UPDATE movies
  set title = $2,
  year = $3,
  runtime = $4,
  genres = $5
WHERE id = $1
`

type UpdateMovieParams struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Year    int32    `json:"year"`
	Runtime Runtime  `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.ExecContext(ctx, updateMovie,
		arg.ID,
		arg.Title,
		arg.Year,
		arg.Runtime,
		pq.Array(arg.Genres),
	)
	return err
}
