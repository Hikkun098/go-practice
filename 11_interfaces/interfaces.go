package main

import "fmt"

type Greeter interface {
	greet() string
}

type Japanese struct {
    name string
}

type American struct {
    name string
}

func (j Japanese) greet() string {
    return fmt.Sprintf("こんにちは、私は%sです", j.name)
}

func (a American) greet() string {
    return fmt.Sprintf("Hello, I'm %s", a.name)
}

func sayHello(g Greeter) {
	fmt.Println(g.greet())
}

func main() {
	j := Japanese{name: "Taro"}
	a := American{name: "John"}

	sayHello(j)
	sayHello(a)
}
