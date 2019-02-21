package main

import "fmt"

type contactInfo struct {
	email string
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
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94800, 
		},
	}
	jim.updateName("Jimmy")
	jim.print()
	// fmt.Printf("%+v", jim)
	// type1
	// alex := person{ firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// type 2
	// var alexander person
	// fmt.Println(alexander)
	// fmt.Printf("%+v\n", alexander)
	
	// type 3
	// var ally person
	// ally.firstName = "Ally"
	// ally.lastName = "Anderson"
	// fmt.Println(ally)
	// fmt.Printf("%+v\n", ally)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}