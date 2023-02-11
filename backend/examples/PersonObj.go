package main

import (
	"fmt"
	"strconv"
)

type PersonStruct struct {
	name          string
	age           int
	gender        string
	favoriteColor string
	cgpa          float32
	grade         string
}

// value from props to parameter
// constructor
func (personPtr *PersonStruct) personConstructor(name string, age int, gender string, favColor string, cgpa float32) {
	personPtr.name = name
	personPtr.age = age
	personPtr.gender = gender
	personPtr.favoriteColor = favColor
	personPtr.cgpa = cgpa
}

func (personStructPtr *PersonStruct) getPersonMthd() {
	fmt.Println("Person name: " + personStructPtr.name + " is " + strconv.Itoa((personStructPtr.age)) + " years old" + " with " + personStructPtr.grade + " grade")
}

func (personPtr *PersonStruct) setNamePersonMthd() {
	personPtr.name = "Hari"
}

func (personPtr *PersonStruct) setAgePersonMthd() {
	personPtr.age = 25
}

func (personPtr *PersonStruct) getGradeMthd() {
	switch {
	case personPtr.cgpa >= 9 && personPtr.cgpa <= 10:
		personPtr.grade = "S"

	case personPtr.cgpa >= 8 && personPtr.cgpa < 9:
		personPtr.grade = "A"

	case personPtr.cgpa >= 7 && personPtr.cgpa < 8:
		personPtr.grade = "B"

	case personPtr.cgpa >= 6 && personPtr.cgpa < 7:
		personPtr.grade = "C"

	case personPtr.cgpa < 6:
		personPtr.grade = "F"

	default:
		fmt.Println("Cannot be greater than 10 or less than 0")

	}
}

func main() {
	var personVh PersonStruct
	personVh.personConstructor("Arvind", 24, "Male", "Blue", 9.19)
	personVh.getPersonMthd()
	personVh.setNamePersonMthd()
	personVh.setAgePersonMthd()
	personVh.getGradeMthd()
	personVh.getPersonMthd()
}
