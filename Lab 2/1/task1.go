package main

import (
	"fmt"
	"crypto/sha256"
	"strconv"
	"strings"
)

type Student struct{
	rollNumber int
	name string
	address string
	subjects []string
}

func NewStudent(rollNo int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.rollNumber = rollNo
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

type StudentList struct{
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string) *Student{
	st := NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)
	return st
}

func main(){
	
	student := new(StudentList)
	course1 := []string{"science", "maths", "physics"}
	course2 := []string{"biology", "calculus", "computers"}
	student.CreateStudent(24, "Armughan", "aaaaaaaaaaa", course1)
	student.CreateStudent(24, "Bob", "bbbbbbbbbb", course2)
	
	// for i := 0; i <= len(student.list)-1; i++{
	// 	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
	// 	fmt.Println("student rollno ", student.list[i].rollNumber)
	// 	fmt.Println("student name ", student.list[i].name)
	// 	fmt.Println("student address ", student.list[i].address)
	// 	fmt.Println("student subjects", student.list[i].subjects)
	// }

	// var hasharr []string
	
	for i := 0; i <= len(student.list)-1; i++{
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("student rollno ", student.list[i].rollNumber)
		fmt.Println("student name ", student.list[i].name)
		fmt.Println("student address ", student.list[i].address)
		fmt.Println("student subjects", student.list[i].subjects)
		var appstring string = ""
		appstring = appstring + strconv.Itoa(student.list[0].rollNumber)
		// fmt.Println(appstring)
		appstring = appstring + student.list[i].name
		appstring = appstring + student.list[i].address

		for j := 0; j <= len(student.list[i].subjects)-1; j++{
			appstring = appstring + student.list[i].subjects[j]
		}
		// fmt.Println(appstring)
		fmt.Printf(fmt.Sprintf("%x", sha256.Sum256([]byte(appstring))))
		fmt.Printf("\n")
	}

}