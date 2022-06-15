package main

import (
	"fmt"
	"strings"
)

type Person struct {
	age       uint8
	lastName  string
	firstName string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode string
}

func main() {
	//
	/*john := Person{age: 33, lastName: "Doe", firstName: "John"}

	fmt.Println(john)

	var mary Person

	mary.age = 25
	mary.lastName = "Oreilly"
	mary.firstName = "Mary"

	fmt.Println(mary)
	fmt.Printf("%+v", mary)*/

	jim := Person{
		age:       22,
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email:   "jim@email.com",
			zipCode: "1234",
		}, // mandatory comma
	}

	// Go uses ByValue as default
	//jimPointer := &jim // gives the memory address of the value this variable is pointing at
	//contactPointer := &jim.contact
	//
	//jimPointer.updateName("Jimmed")
	//contactPointer.updateEmail("Jimmed")
	//jimPointer.print()

	jim.updateName("Jimmed")
	jim.print()

}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person /*Give me the value his memory address is pointing is*/) updateName(newFName string) {
	p.firstName = newFName
	p.contact.email = strings.Replace(p.contact.email, "jim", strings.ToLower(newFName), -1) // Functions
}
