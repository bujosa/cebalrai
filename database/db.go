package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// DB is the database connection
var DB *pgx.Conn

// ConnectDB connects to the database
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = conn
}

// CloseDB closes the database connection
func CloseDB() {
	DB.Close(context.Background())
}