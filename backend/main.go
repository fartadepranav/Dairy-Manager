package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

corsOptions := handlers.AllowedOrigins([]string{"*"})
r := mux.NewRouter()
r.Use(handlers.CORS(corsOptions))
log.Fatal(http.ListenAndServe(":8080", r))

func main() {
	// Create a new Gin router
	router := gin.Default()
	db, err := sql.Open("sqlite3", "test.db")

	fmt.Print(db)
	fmt.Print(err)

	// Customer routes
	router.POST("/api/addCustomer", AddCustomer)

	router.GET("/api/readCustomer", ViewAllCustomers)

	// Run the server on port 8080
	err1 := router.Run(":8080")
	if err1 != nil {
		panic(err1)
	}
}
