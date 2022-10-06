package main

import "fmt"

type Person struct {
	name string
	age int
	job string
	salary int
}

func test(class Person){
	fmt.Println(class.name)
	fmt.Println(class.age)
	fmt.Println(class.job)
	fmt.Println(class.salary)
}

func main(){
	test(Person{name: "armughan", age: 23, job: "engineer", salary: 23000})
}