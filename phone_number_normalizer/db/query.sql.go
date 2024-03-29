// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNumber = `-- name: CreateNumber :one
INSERT INTO NUMBERS (
    input, number
) VALUES (
             $1, $2
         )
RETURNING input, number
`

type CreateNumberParams struct {
	Input  pgtype.Text
	Number string
}

func (q *Queries) CreateNumber(ctx context.Context, arg CreateNumberParams) (Number, error) {
	row := q.db.QueryRow(ctx, createNumber, arg.Input, arg.Number)
	var i Number
	err := row.Scan(&i.Input, &i.Number)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM NUMBERS
WHERE number = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, number string) error {
	_, err := q.db.Exec(ctx, deleteAuthor, number)
	return err
}

const getNumber = `-- name: GetNumber :one
SELECT input, number FROM NUMBERS
WHERE number = $1 LIMIT 1
`

func (q *Queries) GetNumber(ctx context.Context, number string) (Number, error) {
	row := q.db.QueryRow(ctx, getNumber, number)
	var i Number
	err := row.Scan(&i.Input, &i.Number)
	return i, err
}

const listNumbers = `-- name: ListNumbers :many
SELECT input, number FROM NUMBERS
ORDER BY number
limit 10
`

func (q *Queries) ListNumbers(ctx context.Context) ([]Number, error) {
	rows, err := q.db.Query(ctx, listNumbers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Number
	for rows.Next() {
		var i Number
		if err := rows.Scan(&i.Input, &i.Number); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
