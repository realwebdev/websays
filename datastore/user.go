package datastore

import (
	"github.com/realwebdev/clockify/models"
)

func (c *Client) CreateUser(user models.User) error {
	if err := c.Db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUsers() ([]models.User, error) {
	users := []models.User{}
	if err := c.Db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (c *Client) AuthenticateUser(usercred map[string]interface{}) (string, error) {
	user := models.User{}
	if err := c.Db.Table("users").Where(usercred).First(&user).Error; err != nil {
		return "error occured", err
	}

	return user.Username, nil
}

func (c *Client) DeleteUser(user_id uint) error {
	if err := c.Db.Table("users").Where("id = ?", user_id).First(&models.User{}).Error; err != nil {
		return err
	}

	if err := c.Db.Table("users").Where("id = ?", user_id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
