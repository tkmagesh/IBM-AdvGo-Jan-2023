package main

import (
	"log"
	config "todo-app/config"
	models "todo-app/models"
	routes "todo-app/routes"

	_ "net/http/pprof"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildConfig()))

	/*
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	*/

	//to configure the release mode
	// gin.SetMode(gin.ReleaseMode)

	if err != nil {
		log.Fatalf("Error connecting to database : %s", err.Error())
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})
	r := routes.SetupRouter()

	r.Run()
}
