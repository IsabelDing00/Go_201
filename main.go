/*
This is going to show the example of Struct type and Pointer
*/

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // before I have contact and 4th way has contact instead of contanctInfo, but it is ok to have contactInfo struct here like this
}

func main() {
	// TODO 1: 1st way to create a new struct
	// alex := person{"Alex","Anderson"}

	// TODO 2: 2nd way to create a new struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// TODO 3: 3rd way to create a new struct
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// TODO 4: 4th
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 83440,
		},
	}

	// Todo: update the fname, use pass by reference -> Pointer
	jim.updateName("Jimmy")
	jim.print() // use the print() function
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Structs with receiver functions
func (p person) print() {
	fmt.Printf("%+v", p)
}
