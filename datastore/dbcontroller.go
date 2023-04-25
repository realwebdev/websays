package datastore

import (
	"time"

	"github.com/realwebdev/clockify/models"
)

type DBController interface {
	CreateUser(user models.User) error
	GetUsers() (users []models.User, err error)
	AuthenticateUser(usercred map[string]interface{}) (string, error)
	DeleteUser(user_id uint) error

	CreateProject(project models.Project) error
	GetProjects() (projects []models.Project, err error)
	UpdateProject(project_id uint, updates map[string]interface{}) error
	DeleteProject(project_id uint) (string, error)

	StartActivity(activity models.Activity) error
	EndActivity(activity_id uint) (time.Duration, error)
	UpdateActivity(activity_id uint, updates map[string]interface{}) error
	DeleteActivity(activityid uint) error

	AutoMigrate(modelObject interface{}) error

	Close() error
}
