# go-sqlc-experiment

Personal opinion about `sqlc`.

## Pros

* Solid code generation
* Write SQL statements as-is, no need for a complex ORM query builder
* It actually checks if the queries are valid when generating code, which allows for catching errors early in the dev process
* The generated code is straightforward to use
* Works well with existing migration tooling such as [go-migrate](https://github.com/golang-migrate/migrate)
* Documentation is ok, but not that extensive
* Can configure/override parts of the generated code
* Can be used to work with PostGIS to some extent

## Cons

* Sometimes the documentation is a bit lacking
* Does not support PostGIS data types out of the box
* Not possible to create custom types that maps to certain Golang structs
  * For example: `geometry` in PostGIS can be of many types such as a point, line, polygon, etc. We cannot distinguish between these different subtypes at this point of time

## Setup

1. `$ docker compose up`
2. `$ exec dev ./run-test.sh`

## Tips

* Define the type of your field in the query when possible, this will help sqlc infer the correct type for the field. Sometimes it falls back to `interface{}` when not doing so.
* When working with PostGIS data types you can render them as a string and process that separately

## Resources used

* https://zenn.dev/tchssk/articles/a701d3ce5f9b6b
* https://docs.sqlc.dev/en/stable/index.html
* https://play.sqlc.dev/
* https://github.com/kyleconroy/sqlc/issues/231
