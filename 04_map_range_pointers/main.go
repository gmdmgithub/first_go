package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hi there use me")
	mapFunc()
	rangeFunc()
	pointerFunc()
	fmt.Println(time.Since(start))
}

func mapFunc() {
	//map definition
	emails := make(map[string]string)

	//key values
	emails["greg"] = "greg@gmail.com"
	emails["john"] = "john@gmail.com"

	fmt.Println(len(emails))
	fmt.Println("Adam email is ", emails["adam"]) //wrong?
	fmt.Println("John email is ", emails["john"])

	if len(emails["adam"]) == 0 {
		emails["adam"] = "adam@gmail.com"
	}
	fmt.Println(emails)

	//delete one element
	delete(emails, "adam")

	fmt.Println(emails)

	//new map defined at creation
	scores := map[string]int{"adam": 20, "mark": 28, "greg": 60}
	scores["john"] = 11
	//resp, err := http.Get("https//www.google.com")
	http.Get("https//www.google.com")
	fmt.Println(scores)
	log.Fatalf("INFO from LOG %s", scores)

}

func rangeFunc() {
	fmt.Println("##### range func and loops ####")
	someNum := []int{21, 34, 11, 23, 45, 63}

	//range with index
	for i, id := range someNum {
		fmt.Printf("%d - num %d\n", i, id)
	}

	//range withot index
	for id := range someNum {
		fmt.Printf("num %d\n", id)
	}
	//range withot index different way
	for _, id := range someNum {
		fmt.Printf("num with _ %d\n", id)
	}

	fmt.Println(" standard loop ")
	for j := 0; j < len(someNum); j++ {
		fmt.Printf("%d - num %d\n", j, someNum[j])
	}

	//range with map
	scores := map[string]int{"adam": 20, "mark": 28, "greg": 60}
	for key, value := range scores {
		fmt.Printf("Key is %s, and value is %d\n", key, value)
	}

}

func pointerFunc() {
	fmt.Println("##### Pointers ####")
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//change valu with pointer
	*b = 10
	fmt.Println(a)
}
