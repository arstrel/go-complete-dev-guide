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
	// alex := person{firstName: "Alex", lastName: "Anders"}
	// jamie := person{"Jamie", "Jafferson"}
	// var stan person

	// stan.firstName = "Stan"
	// stan.lastName = "Stansfield"

	// fmt.Println(alex, jamie)
	// // print all field names and their values
	// fmt.Printf("%+v\n", stan)

	jim := person{
		firstName: "Jim",
		lastName:  "Jimaloo",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
