package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/config"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/handler"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/router"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/service"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Init()

	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := config.InitializePostgres()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	services := service.NewEmployeeService(db)

	employeeHandler := handler.NewEmployeeHandler(services)

	router.InitializeRoutes(r, employeeHandler)

	errorr := r.Run(":" + port)

	if errorr != nil {
		fmt.Println(errorr)
		return
	}
}
