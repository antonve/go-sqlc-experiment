package main_test

import (
	"context"
	"database/sql"
	"log"

	"github.com/antonve/go-sqlc-experiment/storage"

	"github.com/cridenour/go-postgis"
	_ "github.com/lib/pq"
)

func ExampleSqlc() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "host=postgis user=root dbname=experiment password=hunter2 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	queries := storage.New(db)

	// list all restaurants
	restaurants, err := queries.ListRestaurants(ctx)
	if err != nil {
		log.Fatalf("cannot list all restaurants: %v", err)
	}
	log.Println(restaurants)

	// create a restaurant
	insertedRestaurant, err := queries.CreateRestaurant(ctx, storage.CreateRestaurantParams{
		Name:     "CoCo Ichibanya Ebisu",
		Location: postgis.PointS{SRID: 4326, X: 35.64699825984844, Y: 139.71194575396922},
	})
	if err != nil {
		log.Fatalf("cannot create restaurant: %v", err)
	}
	log.Println(insertedRestaurant)

	// Output:
	// [Book(id=1, name=The Phoenix Project)]
}
