package parser

import "fmt"

func A(fuck int, shit int) {
	cc := 10
	fmt.Println(cc)
}

type Person struct {
	name string
	age  int32
}

func (p Person) IsAdult() bool {
	return p.age >= 18
}

func (p Person) Eat(food string, amount int) string {
	fmt.Println("eat eat")
	return ""
}
