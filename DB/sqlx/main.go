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
type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) getEmployeeRecords() {
	rows, err := r.db.Query("select * from employees")
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
}
func (r *Repo) getSalaryRecords() {
	rows, err := r.db.Query("select * from salaries")
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
}
func (r *Repo) addEmployee() {
	fmt.Println("Please enter the name of a new employee")
	var input string
	fmt.Scan(&input)
	sqlStatement := `INSERT INTO employees (name) VALUES ($1)`
	_, err := r.db.Exec(sqlStatement, input)
	if err != nil {
		panic(err)
	}
}
func (r *Repo) addSalary() {
	fmt.Println("Please enter the name and employee`s id")
	var name, emp_id string
	fmt.Scan(&name, &emp_id)
	sqlStatement := `INSERT INTO salaries (employee_id, name) VALUES ($1, $2)`

	rez, err := r.db.Query(fmt.Sprintf("SELECT * FROM employees WHERE id=%s", emp_id))
	if rez != nil {
		_, err = r.db.Exec(sqlStatement, emp_id, name)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sqlx.Open("postgres", "user=postgres dbname=taskDb password=12345 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	repo := New(db)

	repo.getEmployeeRecords()
	repo.getSalaryRecords()
	repo.addEmployee()
	repo.addSalary()
}
