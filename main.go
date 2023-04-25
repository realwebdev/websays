package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/clockify/conf"
	"github.com/realwebdev/clockify/datastore"
	"github.com/realwebdev/clockify/handlers"
	"github.com/realwebdev/clockify/routers"
)

func main() {
	conf := conf.New()
	dbhandler, err := datastore.New(conf.ConnStr())
	if err != nil {
		log.Fatal("error occured", err.Error())
	}
	defer dbhandler.Close()
	// dbhandler.AutoMigrate(models.Activity{})
	// dbhandler.AutoMigrate(models.User{})
	// dbhandler.AutoMigrate(models.Project{})
	// dbhandler.AutoMigrate(models.Authentication{})
	// dbhandler.AutoMigrate(models.Token{})

	controller := handlers.New(dbhandler)
	r := routers.SetupRouter(controller)
	checkServer := r.Run(":8080")
	if checkServer != nil {
		log.Fatal("Failed to start server: ", err)
	}

	// models.SignUp(models.User{Username: "haseeb", Email: "saskitchawn@gmail.com3.3", Password: "mypass2"}, db)
	// models.SignIn("saskitchawn@gmail.com1", "mypass1", db)
	// models.UserDeletion(1, db)

	// models.CreateProject(models.Project{UserID: 2, Project_name: "MYFirstProj3"}, db)
	// models.DeleteProject(2, db)
	// updates := make(map[string]interface{})
	// updates["project_name"] = "mynewname"
	// models.UpdateProject(1, updates, db)

	// models.CreateStartActivity(models.Activity{Activity_name: "MYFirsActivity2", ProjectID: 1, UserID: 1, Start_time: time.Now()}, db)
	// models.EndActivity(1, db)
	// models.DeleteActivity(4, db)
	// models.UpdateActivity(1, "MYFirstActivityChanged", db)

	// check category and article using postman

}
