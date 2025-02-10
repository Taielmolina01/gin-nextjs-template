package application

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

const (
	MESSAGE_ERROR_CONNECTING_DB     = "Error connecting to the database"
	MESSAGE_ERROR_CREATING_DB       = "Database already exists or had ocurred an error creating it"
	MESSAGE_DATABASE_ALREADY_EXISTS = "Database already exists"
	MESSAGE_SUCCESS_CREATING_DB     = "Success creating the database"
	MESSAGE_SUCCESS_CREATING_TABLES = "Success creating all the tables of the database"
	MESSAGE_ERROR_CREATING_TABLE    = "Error creating table: %w"
	DATABASE_NAME                   = "db-gin-template"
)

func ConnectDB(config *Configuration) (*gorm.DB, error) {
	dsn := config.DbDsn

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
	}

	createDatabaseIfNotExist(dsn)

	db, err = gorm.Open(postgres.Open(dsn+" dbname="+DATABASE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
		return nil, err
	}

	tables := GetAllModels()

	for _, t := range tables {
		err = db.AutoMigrate(t)

		if err != nil {
			return nil, fmt.Errorf(MESSAGE_ERROR_CREATING_TABLE, err)
		}
	}

	fmt.Println(MESSAGE_SUCCESS_CREATING_TABLES)

	return db, nil
}

func createDatabaseIfNotExist(dsn string) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(MESSAGE_ERROR_CONNECTING_DB, err)
	}

	var exists bool
	conn.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", DATABASE_NAME).Scan(&exists)

	if exists {
		fmt.Println(MESSAGE_DATABASE_ALREADY_EXISTS)
		return
	}

	err = conn.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, DATABASE_NAME)).Error
	if err != nil {
		fmt.Println(MESSAGE_ERROR_CREATING_DB, err)
	} else {
		fmt.Println(MESSAGE_SUCCESS_CREATING_DB)
	}
}
