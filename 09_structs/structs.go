package main

import "fmt"

type Person struct {
	name      string
	age       int
	job       string
	country   string
	working   bool
	languages []string
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old. I work as a %s in %s.\n", p.name, p.age, p.job, p.country)
	if p.working {
		fmt.Println("I am currently working.")
	} else {
		fmt.Println("I am currently not working.")
	}
	fmt.Printf("I speak the following languages: %v\n", p.languages)
}

func main() {
	person := Person{
		name:      "Suzuki Taro",
		age:       30,
		job:       "Software Engineer",
		country:   "Japan",
		working:   true,
		languages: []string{"Japanese", "English"},
	}

	person.greet()
}
