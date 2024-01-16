package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var blah = "blah"

func main() {
	// Create a new Gin router
	router := gin.Default()
	db, err := sql.Open("sqlite3", "test.db")

	fmt.Print(db)
	fmt.Print(err)

	// Define your routes and handlers
	//router.GET("/api/test", Testget)

	//router.GET("/api/getData", Readdata)

	//router.POST("/api/post", Posthandler)

	// Customer routes
	router.POST("/api/addCustomer", AddCustomer)

	router.GET("/api/readCustomer", ViewAllCustomers)

	// Run the server on port 8080
	err1 := router.Run(":8080")
	if err1 != nil {
		panic(err1)
	}
}
