-- name: ListRestaurants :many
select *
from restaurants
order by name;

-- name: CreateRestaurant :one
insert into restaurants (
  name, location
) values (
  $1, $2
)
returning $1;
