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
	Config.DB.AutoMigrate(&Models.Catergory{})

	r := Routes.SetupRouter()

	r.LoadHTMLFiles(
		"UI/Templates/index.html",
		"UI/Templates/Pages/Users/userDashboard.tmpl",
		"UI/Templates/Pages/Admin/adminDashboard.tmpl",
		"UI/Templates/Pages/Admin/brands.tmpl",
		"UI/Templates/Pages/Admin/products.tmpl",
		"UI/Templates/Layouts/AdminLayout.tmpl",
		"UI/Templates/Layouts/Admin/sidebar.tmpl",
		"UI/Templates/Layouts/Admin/header.tmpl",

		//Customer Pages
		//"UI/Templates/Pages/Customer/CustomerLanding.tmpl",
	)

	/*r.Static("/UI/Statics/CSS", "adminStyle.css")
	r.Static("/UI/Statics/JS", "admin.js") */
	//r.Static("/UI/Statics/", "./assets")

	r.Run()
}
