package main

import (
	"ProjectCinema/config"
	"ProjectCinema/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Hello Cinema")
	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	_ = r.Run("localhost:8086")

}
