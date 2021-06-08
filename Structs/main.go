package main

import "fmt"

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

	// 1. basic way to use struct

	//Hjani := person{firstName: "Hardik", lastName: "Jani"}
	//fmt.Println(Hjani)

	//2. one more to write struct

	// var hjani person

	// hjani.firstName = "Hardik"
	// hjani.lastName = "Jani"

	// fmt.Println(hjani)
	// fmt.Printf("%+v", hjani)

	//3. Embedding structs

	hardik := person{
		firstName: "Hardik",
		lastName:  "Jani",
		contactInfo: contactInfo{
			email:   "hjani@gmail.com",
			zipCode: 12345,
		},
	}

	// struct with pointers
	hjaniPointer := &hardik
	hjaniPointer.updateName("hardik68")
	hardik.Print()

}

// struct with pointers
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Pass by value
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

//structs with receiver function
func (p person) Print() {
	fmt.Printf("%+v", p)
}
