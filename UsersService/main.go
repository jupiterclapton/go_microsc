package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cache := Cache{Enable: true}
	flag.StringVar(
		&cache.Address,
		"redis_address",
		os.Getenv("APP_RD_ADDRESS"),
		"Redis Address",
	)

	flag.StringVar(
		&cache.Auth,
		"redis_auth",
		os.Getenv("APP_RD_AUTH"),
		"Redis Auth",
	)

	flag.StringVar(
		&cache.DB,
		"redis_db_name",
		os.Getenv("APP_RD_DBNAME"),
		"Redis DB name",
	)

	flag.IntVar(
		&cache.MaxIdle,
		"redis_max_idle",
		10,
		"Redis Max Idle",
	)

	flag.IntVar(
		&cache.MaxActive,
		"redis_max_active",
		100,
		"Redis Max Active",
	)

	flag.IntVar(
		&cache.IdleTimeoutSecs,
		"redis_timeout",
		60,
		"Redis timeout in seconds",
	)
	flag.Parse()
	cache.Pool = cache.NewCachePool()
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a := App{}
	a.Initialize(cache, db)
	a.Run(":8080")
}
