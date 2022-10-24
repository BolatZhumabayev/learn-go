package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	user1 := person{firstName: "Alex", lastName: "Smirnoff",
		contact: contactInfo{
			email:   "max.smirnoff@gmail.com",
			zipCode: 1292},
	}
	user1.print()

	var user2 person
	user2.firstName = "Max"
	user2.lastName = "Anderson"

	user2.updateName("Maximus")

	user2.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	fmt.Println(p)
	fmt.Println(*p)
	(*p).firstName = newFirstName
}
