package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	//access the database
	model.InitDb()

	routes.InitRouter()
}
