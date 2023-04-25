package datastore

import (
	"time"

	"github.com/realwebdev/clockify/models"
)

func (c *Client) StartActivity(activity models.Activity) error {
	starttime := time.Now()
	activity.Start_time = starttime
	if err := c.Db.Create(&activity).Error; err != nil {
		return err
	}

	return nil
}

func (c *Client) EndActivity(activity_id uint) (time.Duration, error) {
	activity := models.Activity{}
	if err := c.Db.Table("activities").Where("ID = ?", activity_id).First(&activity).Error; err != nil {
		return 0, err
	}

	starttime := activity.Start_time
	endTime := time.Now()
	difference := endTime.Sub(starttime)
	if err := c.Db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"total_time": difference, "end_time": endTime}).Error; err != nil {
		return 0, err
	}

	return difference, nil
}

func (c *Client) UpdateActivity(activity_id uint, updates map[string]interface{}) error {
	if err := c.Db.Table("activities").Where("ID = ?", activity_id).First(&models.Activity{}).Error; err != nil {
		return err
	}

	if err := c.Db.Table("activities").Where("ID = ?", activity_id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteActivity(activityid uint) error {
	if err := c.Db.Table("activities").Where("ID = ?", activityid).First(&models.Activity{}).Error; err != nil {
		return err
	}

	if err := c.Db.Table("activities").Where("ID=?", activityid).Delete(&models.Activity{}).Error; err != nil {
		return err
	}

	return nil
}
