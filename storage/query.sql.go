// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package database

import (
	"context"
)

const createRestaurant = `-- name: CreateRestaurant :one
insert into restaurants (
  name, location
) values (
  $1, $2
)
returning $1
`

type CreateRestaurantParams struct {
	Name     string
	Location interface{}
}

func (q *Queries) CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, createRestaurant, arg.Name, arg.Location)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getRestaurants = `-- name: GetRestaurants :many
select id, name, location
from restaurants
order by name
`

func (q *Queries) GetRestaurants(ctx context.Context) ([]Restaurant, error) {
	rows, err := q.db.QueryContext(ctx, getRestaurants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Restaurant
	for rows.Next() {
		var i Restaurant
		if err := rows.Scan(&i.ID, &i.Name, &i.Location); err != nil {
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
