package main

import (
	"database/sql"
	"fmt"
	"os"
)

func initializeDB() (*sql.DB, error) {
	//dbURL := os.Getenv("DATABASE_URL") // Make sure this environment variable is set in Render's settings
	//if dbURL == "" {
	//	log.Fatal("DATABASE_URL is not set")
	//}
	//
	//db, err := sql.Open("postgres", dbURL)
	//if err != nil {
	//	return nil, err
	//}
	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}
	//
	////migrationUp(db)
	//
	//return db, nil
	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("user"),
		os.Getenv("password"), os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//migrationUp(db)

	return db, nil
}
