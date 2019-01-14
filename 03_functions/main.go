package main

import (
	"fmt"
)

func saySomething(text string) string {
	return "Reurn some text with:  " + text
}

func sayMore(text string, age int) (string, int) {
	return " More than this " + text, age + 2
}

func arrayExample() {

	var ages [2]int
	ages[0] = 37
	ages[1] = 35
	fmt.Println("Age of my fiend is %d", ages)

	//slices
	nameSlice := []string{"Greg", "Art", "Criss", "John", "Mark"}

	fmt.Println(len(nameSlice))
	fmt.Println(nameSlice[2:4])
}

func tagFunc(tag string) func(string) string {
	tagM := func(mes string) string {
		return "<" + tag + ">" + mes + "</" + tag + ">"
	}
	return tagM
}

func main() {

	fmt.Println("Functions in action!")

	fmt.Println(saySomething("Greg"))
	fmt.Println(sayMore("Grer", 37))

	arrayExample()
	h1 := tagFunc("H1")
	fmt.Println(h1("Main function call"))
}
