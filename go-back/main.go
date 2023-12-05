package main

import (
	"cyclesaster/data_handling"
	"cyclesaster/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	data_handling.Load_data("../Accidents/Accidents2005-2022")

	fmt.Println("Booting up server...")

	routers.SetupRouter(router)

	router.Run(":8080")
}
