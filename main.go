package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	condb, errdb := sql.Open("mssql", "server=xxxxxxxxx;user id=xxxxx;password=xxxx;database=xxxx;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	var (
		id           int
		FirstName    string
		LastName     string
		Department   string
		EmployeeType string
	)

	rows, err := condb.Query("Select EmployeeID,Firstname,Lastname,Dpartment,Employeetype from employees")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Employee ID, First Name, Last Name, Department, EmployeeType")
	for rows.Next() {
		err := rows.Scan(&id, &FirstName, &LastName, &Department, &EmployeeType)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(id, FirstName, LastName, Department, EmployeeType)
	}
	defer condb.Close()

}
