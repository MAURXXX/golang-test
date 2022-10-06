package main

import (
	"fmt"
	"strings"
)

type Student struct{
	rollNumber int
	name string
	address string
}

func NewStudent(rollNo int, name string, address string) *Student {
	s := new(Student)
	s.rollNumber = rollNo
	s.name = name
	s.address = address
	return s
}

type StudentList struct{
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student{
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main(){
	
	student := new(StudentList)
	student.CreateStudent(24, "Armughan", "aaaaaaaaaaa")
	student.CreateStudent(24, "Bob", "bbbbbbbbbb")
	
	for i := 0; i <= len(student.list)-1; i++{
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("student rollno ", student.list[i].rollNumber)
		fmt.Println("student name ", student.list[i].name)
		fmt.Println("student address ", student.list[i].address)
	}
		

}