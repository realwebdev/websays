package datastore

import (
	"log"
)

func (c *Client) AutoMigrate(modelObject interface{}) error {
	// if err := c.Db.DropTableIfExists(modelObject).Error; err != nil {
	// 	log.Print("Error occured While Deleting Previous Table")
	// 	return err
	// }
	// log.Print("Previous Table drop successfully")

	if err := c.Db.AutoMigrate(modelObject); err != nil {
		return err
	}
	log.Print("models created successfully")

	return nil
}
