package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBInstance() *sql.DB {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file.")
	}

	connectionString :=
		os.Getenv("DB_USER") + os.Getenv("DB_PASSWORD") + ":@" + "(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
