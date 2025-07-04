package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//Refer to : https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797350#questions
	//another long way to do it is jimPointerToPerson = &jim then jimPointerToPerson.updateName("jimmy")
	//because its supposed to be a pointer, but go provides this shortcut
	jim.updateName("jimmy")
	jim.print()
	// runMapExample()
	//runChannels()
	//db.RunDB()
	//db.Connect()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
