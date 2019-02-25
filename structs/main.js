let person = {
		firstName: "Jim",
		lastName: "Party",
}

console.log(person)

const updateFirstName = firstname => person.firstName = firstname


updateFirstName("Jimmy")
console.log(person)