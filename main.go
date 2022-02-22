package main

import (
	"fmt"

	"alexedwards.net/CC-GO/Config"
	"alexedwards.net/CC-GO/Models"
	"alexedwards.net/CC-GO/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	fmt.Println(Config.DB)
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Todo{})
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()

	/* r.LoadHTMLGlob("UI/Templates/*.html")
	r.LoadHTMLGlob("UI/Templates/Pages/Users/*.html")
	r.LoadHTMLGlob("UI/Templates/Pages/Admin/*.html") */
	r.LoadHTMLFiles(
		"UI/Templates/index.html",
		"UI/Templates/Pages/Users/userDashboard.html",
		"UI/Templates/Pages/Admin/adminDashboard.html")

	r.Run()
}
