package main

import (
	"fmt"
	"strconv"
)

func saySomething(text string) string {
	return "Rerun some text with:  " + text
}

func sayMore(text string, age int) (string, int) {
	return " More than this " + text, age + 2
}

func arrayExample() {

	var ages [2]int
	ages[0] = 37
	ages[1] = 35
	fmt.Printf("Age of my fiend is %d\n", ages)

	//slices
	nameSlice := []string{"Greg", "Art", "Criss", "John", "Mark"}

	fmt.Println(len(nameSlice))
	fmt.Println(nameSlice[2:4])
	i := 0
	for i < len(nameSlice) {
		fmt.Println(nameSlice[i])
		i++
	}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}
}

func tagFunc(tag string) func(string) string {
	tagM := func(mes string) string {
		return "<" + tag + ">" + mes + "</" + tag + ">"
	}
	return tagM
}

func checkColour(colour string) {
	switch colour {
	case "red":
		fmt.Println("colour is red")
	case "green":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is red")
	default:
		fmt.Println("you send wrong colour")
	}
}

// EXCEPTING INPUT
func getInput() {
	fmt.Println("###### working with user input ######")

	fmt.Println("What is your name? ")

	var name string

	fmt.Scan(&name)

	fmt.Println("Hello", name)

}

func main() {

	fmt.Println("Functions in action!")

	fmt.Println(saySomething("Greg"))
	fmt.Println(sayMore("Greg", 37))

	arrayExample()
	h1 := tagFunc("H1")
	fmt.Println(h1("Main function call"))

	checkColour("green")
	strincFunc()

	fileOperation()
	fileStatistics()
	checkSum()

	getInput()
	conversingCasting()

}

// CASTING AND CONVERTING
func conversingCasting() {

	fmt.Println("###### Casting and converting ######")

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	// Convert numbers types
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	// Convert a string into an int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	// Convert a string into a float
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)
}
