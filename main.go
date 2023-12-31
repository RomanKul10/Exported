package main

import (
	"log"
	"myapp/staff"
)

var employee = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employee,
	}

	//myStaff.All()

	//log.Println(myStaff.All())
	//staff.OverpaidLimit = 10000
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff", myStaff.Underpaid())

}
