package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	var name = "Greg"
	var age = 37
	tall, surname := 1.85, "Doe"
	//surename := "Doe"
	const isHight = true //cannot be redefined
	fmt.Printf("Hello World "+name+" in age %d is height %t and surename is %s and tall %f   \n", age, isHight, surname, tall)
	fmt.Printf("%T\n", isHight)
	fmt.Printf("Some math %v\n", math.Floor(23.45))
	fmt.Printf("Some other %v\n", rand.Intn(100)) //81 withoud a seed

	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Printf("Some other %v\n", rand.Intn(100))
}
