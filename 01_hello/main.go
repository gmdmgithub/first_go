package main

import "fmt"

func main() {

	var name = "Greg"
	var age = 37
	tall, surname := 1.85, "Doe"
	//surename := "Doe"
	const isHight = true //cannot be redefined
	fmt.Printf("Hello World "+name+" in age %d is height %t and surename is %s and tall %f   \n", age, isHight, surname, tall)
	fmt.Printf("%T\n", isHight)
}
