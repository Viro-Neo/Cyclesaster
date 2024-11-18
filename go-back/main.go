package main

import (
	"cyclesaster/database"
	"cyclesaster/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.InitDB()
	defer db.Close()

	//data_handling.Load_data(db, "../Accidents/Accidents2005-2022")

	fmt.Println("Booting up server...")

	routers.SetupRouter(db, router)

	router.Run(":8080")
}
