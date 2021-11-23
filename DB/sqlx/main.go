package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type employee struct {
	id   int
	name string
}
type salary struct {
	id     int
	id_emp int
	name   string
}

func getEmployees() {

}

func main() {
	db, err := sqlx.Open("postgres", "user=postgres dbname=taskDb password=12345 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//Retrieve the list of employees
	rows, err := db.Query("select * from employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	employees := []employee{}

	for rows.Next() {
		p := employee{}
		err := rows.Scan(&p.id, &p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, p)
	}
	for _, p := range employees {
		fmt.Println(p.id, p.name)
	}
	//Retrieve the list of salary records
	rows, err = db.Query("select * from salary")
	if err != nil {
		err.Error()
	}
	defer rows.Close()
	salaries := []salary{}

	for rows.Next() {
		p := salary{}
		err := rows.Scan(&p.id, &p.id_emp, &p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		salaries = append(salaries, p)
	}
	for _, p := range salaries {
		fmt.Println(p.id, p.id_emp, p.name)
	}

	//Add employee
	fmt.Println("Please enter the name of a new employee")
	var input string
	fmt.Scan(&input)
	sqlStatement := `INSERT INTO employee (name) VALUES ($1)`
	_, err = db.Exec(sqlStatement, input)
	if err != nil {
		panic(err)
	}

	//Add salary record
	fmt.Println("Please enter the name and employee`s id")
	var name, emp_id string
	fmt.Scan(&name, &emp_id)
	sqlStatement := `INSERT INTO salary (employee_id, name) VALUES ($1, $2)`

	rez, err := db.Query(fmt.Sprintf("SELECT * FROM employee WHERE id=%s", emp_id))
	if rez != nil {
		_, err = db.Exec(sqlStatement, emp_id, name)
		if err != nil {
			panic(err)
		}
	}

}
