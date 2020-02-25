package parser

import "fmt"

func NormalFunc(param1 int, param2 int) {
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
