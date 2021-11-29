package main

import (
	"bufio"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Employee struct {
	Id   int
	Name string
}
type Salary struct {
	Id     int
	Id_emp int `gorm:"column:employee_id"`
	Name   string
}

func main() {
	dsn := "user=postgres password=12345 dbname=taskDb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = db

	//read the employees
	var employees []*Employee
	res := db.Find(&employees)
	if res.Error != nil {
		fmt.Println("Could not find")
	}
	for _, employee := range employees {
		fmt.Println(employee)
	}
	//read salary
	var salaries []*Salary
	res = db.Find(&salaries)
	if res.Error != nil {
		fmt.Println("Could not find")
	}
	for _, salary := range salaries {
		fmt.Println(salary)
	}

	var scanner = bufio.NewScanner(os.Stdin)
	//add employee
	fmt.Println("Please enter the name of a new employee")
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	toSave := &Employee{Name: scanner.Text()}
	db.Save(toSave)

	//add salary
	var toAddSalary = &Salary{}
	fmt.Println("Please enter the count and employee`s id")
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	toAddSalary.Name = scanner.Text()
	fmt.Println("Please enter employee`s id")
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	toAddSalary.Id_emp, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	var toAddemployee = &Employee{}
	db.First(&toAddemployee, "id = ?", toAddSalary.Id_emp)
	if toAddemployee.Name != "" {
		db.Save(toAddSalary)
	}

}
