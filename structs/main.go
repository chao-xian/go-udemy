package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email    string
	postCode string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:    "jim@jimail.com",
			postCode: "E1 7AB",
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// var alex person
// alex.firstName = "Alex"
// fmt.Println(alex)
// alex := person{firstName: "Alex", lastName: "Anderson"}

// 	fmt.Printf("%+v", alex)

// contact: contactInfo

// jimPointer := &jim
// jimPointer.updateName("Jimmy")

// pointer = &variable - "Give me the memory address that variable is at"
// *pointer "Give me the value of what is at that address"
// *type "A pointer that points at a type" a Type Description
// with a receiver in func ...(... *type) { ...*pointer... } Go shortcuts and turns the variable pointer behind the scenes
