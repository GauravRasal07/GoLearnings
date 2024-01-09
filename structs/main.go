package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email  string
	mobile int
}

func (pointerToPerson *person) updateFirstname(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

func main() {
	// gaurav := person{"Gaurav", "Rasal"}

	// gaurav := person{firstName: "Gaurav", lastName: "Rasal"}

	// var gaurav person
	// gaurav.firstName = "Gaurav"
	// gaurav.lastName = "Rasal"

	gaurav := person{
		firstName: "Gaurav",
		lastName:  "Rasal",
		contactInfo: contactInfo{
			email:  "gaurav.rasal@gmail.com",
			mobile: 9876543210,
		},
	}

	// pointer := &gaurav				// First approach
	// pointer.updateFirstname("Mansi")
	gaurav.updateFirstname("Mansi")		// Shortcut for pointer

	// fmt.Println(gaurav)
	// fmt.Printf("%+v", gaurav)
	gaurav.printPerson()
}
