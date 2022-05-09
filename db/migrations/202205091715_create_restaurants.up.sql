create table restaurants (
  id bigserial primary key,
  name text not null,
  -- use wgs84
  location geometry(point, 4326) not null
);

create index idx_restaurants_location
  on restaurants using gist(location);
