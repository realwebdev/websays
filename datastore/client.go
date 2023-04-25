package datastore

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	Db *gorm.DB
}

func New(connStr string) (h DBController, err error) {
	db, err := ConnectDB(connStr)
	if err != nil {
		return nil, err
	}

	return &Client{Db: db}, nil
}

func (c *Client) Close() error {
	db, err := c.Db.DB()
	if err != nil {
		fmt.Println("error occured")
	}

	return db.Close()
}

func ConnectDB(dbURI string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Database!")

	return db, nil
}
