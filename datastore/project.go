package datastore

import (
	"github.com/realwebdev/clockify/models"
)

func (c *Client) CreateProject(project models.Project) error {
	if err := c.Db.Create(&project).Error; err != nil {
		return err
	}

	return nil
}

func (c *Client) GetProjects() (projects []models.Project, err error) {
	if err := c.Db.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *Client) UpdateProject(project_id uint, updates map[string]interface{}) error {
	if err := c.Db.Table("projects").Where("ID = ?", project_id).First(&models.Project{}).Error; err != nil {
		return err
	}

	if err := c.Db.Table("projects").Where("ID = ?", project_id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteProject(project_id uint) (string, error) {
	project := models.Project{}
	if err := c.Db.Table("projects").Where("ID = ?", project_id).First(&project).Error; err != nil {
		return "error occured", err
	}

	if err := c.Db.Table("projects").Where("ID = ?", project_id).Delete(&project).Error; err != nil {
		return "error occured", err
	}

	return project.Project_name, nil
}
