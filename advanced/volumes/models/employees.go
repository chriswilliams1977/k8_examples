package models

import (
	"fmt"
)

type Employee struct {
	Id int32
	Firstname   string
	Lastname  string
}

func (db *DB) AllEmployees() ([]*Employee, error) {

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	defer rows.Close()

	employees := make([]*Employee, 0)
	for rows.Next() {
		em := new(Employee)
		err := rows.Scan(&em.Id, &em.Firstname, &em.Lastname)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
		employees = append(employees, em)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return employees, nil
}