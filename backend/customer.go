package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

Schema for Customers:
-> Id - Int
-> Name - String
-> Phone - String
-> Address - String

*/

// CREATE TABLE IF NOT EXISTS Customer(id INTEGER PRIMARY KEY, name TEXT, phone TEXT, address TEXT);

func AddCustomer(c *gin.Context) {
	// Get data from the request, for example, the JSON payload
	db, err3 := sql.Open("sqlite3", "dairy.db")
	fmt.Print(err3)
	create := `CREATE TABLE IF NOT EXISTS Customer(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, Phone TEXT, Address TEXT);`
	_, err := db.Exec(create)

	var customer struct {
		Name    string
		Phone   string
		Address string
	}

	c.Bind(&customer)

	fmt.Print(customer)

	c.JSON(200, gin.H{
		"message": "New Customer added successfully: " + customer.Name,
		"data":    customer.Name,
	})

	// Prepare statement for inserting data

	stmt, err := db.Prepare("INSERT INTO Customer(Name,Phone,Address) VALUES( ? ,? ,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the prepared statement
	_, err = stmt.Exec(customer.Name, customer.Phone, customer.Address)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted successfully")

}

func ViewAllCustomers(c *gin.Context) {
	db, err1 := sql.Open("sqlite3", "dairy.db")
	data, err := db.Query("SELECT * FROM Customer;")

	type Customer struct {
		Id      int
		Name    string
		Phone   string
		Address string
	}

	var customerList []Customer = make([]Customer, 0)

	for data.Next() {

		var singleCustomer Customer

		// Scan the columns into variables
		err := data.Scan(&singleCustomer.Id, &singleCustomer.Name, &singleCustomer.Phone, &singleCustomer.Address)
		if err != nil {
			fmt.Print(err)
		}

		// Print or process the extracted data
		customerList = append(customerList, singleCustomer)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All customers",
		"data":    customerList,
	})
	fmt.Print(err1)
	fmt.Print(err)
}
