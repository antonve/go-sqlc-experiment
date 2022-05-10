package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/antonve/go-sqlc-experiment/db"

	"github.com/cridenour/go-postgis"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ExampleSqlc() {
	ctx := context.Background()

	psql, err := sql.Open("pgx", "host=postgis user=root dbname=experiment password=hunter2 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	runMigrations(psql)

	queries := db.New(psql)

	// create restaurants
	cocoId := createRestaurant(queries, "CoCo Ichibanya Ebisu", 35.64699825984844, 139.71194575396922)

	// list all restaurants
	restaurants, err := queries.ListRestaurants(ctx)
	if err != nil {
		log.Fatalf("cannot list all restaurants: %v", err)
	}
	fmt.Println(cocoId, restaurants)

	// Output:
	// 1 [{1 CoCo Ichibanya Ebisu 0101000020E6100000194FC7D6D0D241405D3A7642C8766140}]
}

func createRestaurant(queries *db.Queries, name string, long, lat float64) int64 {
	id, err := queries.CreateRestaurant(context.Background(), db.CreateRestaurantParams{
		Name:     name,
		Location: postgis.PointS{SRID: 4326, X: long, Y: lat},
	})
	if err != nil {
		log.Fatalf("cannot create restaurant: %v", err)
	}
	return id
}

func runMigrations(psql *sql.DB) {
	driver, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed getting postgres instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed initializing migrations: %v", err)
	}

	err = m.Down()
	if err != nil {
		log.Fatalf("failed dropping database: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}
}
