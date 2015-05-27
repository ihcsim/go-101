package main

import "fmt"

func main() {
	firstPerson := Person{"Ivan Sim", "ihcsim@gmail.com", 35}

	var stringer interface{}
	stringer = firstPerson
	switch str := stringer.(type) {
	case fmt.Stringer:
		fmt.Println(str.String())
	case fmt.GoStringer:
		fmt.Println(str.GoString())
	}
}

type Person struct {
	firstName string
	email     string
	age       int
}

func (person Person) String() string {
	return fmt.Sprintf("%s, %s, %d\n", person.firstName, person.email, person.age)
}

func (person Person) GoString() string {
	return fmt.Sprintf("{First name: %s, Email: %s, Age: %d}", person.firstName, person.email, person.age)
}
