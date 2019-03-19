package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Person - simple person struct
type Person struct {
	firstName, lastName string
	city                string
	gender              string
	age                 int
}

//value reciever
func (p Person) sayHello() string {
	return "Hi there it's " + p.firstName + " and my age is " + strconv.Itoa(p.age)
}

// pointer reciever
func (p *Person) addYear() {
	p.age++
}

// Vehicle - define interface with method
type Vehicle interface {
	weelNumber() int
}

// Car - simple car struct
type Car struct {
	name  string
	weels int
}

// Bike - simple bike struct
type Bike struct {
	weels   int
	name    string
	handars bool
}

func (c Car) weelNumber() int {
	return c.weels
}

func (b Bike) weelNumber() int {
	return b.weels
}

func getWeels(v Vehicle) int {
	return v.weelNumber()
}

// explain - shows the strength of interface and dynamic type
func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Printf("i stored something else type: %T\n", i)
	}
}

func main() {
	fmt.Println("Hi there!")
	person := Person{firstName: "Anna", lastName: "Smith", city: "NY", gender: "f", age: 27}
	person2 := Person{"Adam", "Jordan", "LA", "m", 43}
	fmt.Println(person)
	fmt.Println(person2)

	fmt.Println(person.city)
	person2.age++
	fmt.Println(person2)

	fmt.Println(person2.sayHello())
	person.addYear()
	fmt.Printf("Person %v\n", person)
	fmt.Printf("Person with keys %+v\n", person)
	fmt.Printf("Person representation %#v\n", person)
	fmt.Printf("Person representation city %#v\n", person.city)
	fmt.Printf("Person type %T\n", person)

	volvo := Car{"VXC", 4}
	mtb := Bike{2, "Bike", true}

	fmt.Printf("Car has %v weels\n", getWeels(volvo))
	fmt.Printf("Bike has %d weels\n", getWeels(mtb))

	explain("Hello World")
	explain(52)
	explain(true)
	explain(43.2)
	explain(3 + 2i)
}
