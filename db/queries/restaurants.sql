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

-- name: ListRestaurantsNearby :many
select
  id,
  name,
  GeomFromEWKB(location) as location
from restaurants
where
  ST_Distance(location, GeomFromEWKB(sqlc.arg('origin'))::geometry, false) < sqlc.arg('max_distance')::int
order by id asc;
