package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName string
	city                string
	gender              string
	age                 int
}

//value reciever
func (person Person) sayHello() string {
	return "Hi there it's " + person.firstName + " and my age is " + strconv.Itoa(person.age)
}

//pointer reciever
func (p *Person) addYear() {
	p.age++
}

//interface
type Vehicle interface {
	nweels() int
}

type Car struct {
	name  string
	weels int
}

type Bike struct {
	weels   int
	name    string
	handars bool
}

func (c Car) nweels() int {
	return c.weels
}

func (b Bike) nweels() int {
	return b.weels
}

func getWeels(v Vehicle) int {
	return v.nweels()
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
}
