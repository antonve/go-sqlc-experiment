-- name: ListRestaurants :many
select
  id,
  name,
  GeomFromEWKB(location) as location
from restaurants
order by name;

-- name: CreateRestaurant :one
insert into restaurants (
  name, location
) values (
  sqlc.arg('name'), GeomFromEWKB(sqlc.arg('location'))
)
returning id;
