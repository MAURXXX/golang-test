package main

import "fmt"

type employee struct {
	name string
	salary int
	position string
}
	
type company struct {
	companyName string
	employees []employee
}

func test1(class company){
	fmt.Println(class.companyName)
	fmt.Println(class.employees)
}

func main(){
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Liaquat", 90000, "Half-Stack Developer"}
	emp3 := employee{"Khan", 100000, "Lower Half-Stack Developer"}
	
	cObj := company{
		companyName: "bhaiKehChaarDost",
		employees: []employee{emp1, emp2, emp3},
	}

	fmt.Println(cObj)
}