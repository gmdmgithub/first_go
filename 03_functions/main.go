package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {

	fmt.Println("Functions in action!")

	fmt.Println(saySomething("Greg"))
	fmt.Println(sayMore("Grer", 37))

	arrayExample()
	h1 := tagFunc("H1")
	fmt.Println(h1("Main function call"))

	checkColour("green")
	strincFunc()
	fileOperation()
	getInput()
	conversingCasting()

}

// STRING FUNCTIONS
func strincFunc() {
	fmt.Println("###### Strings operations ######")
	//string for manipulation
	sampString := "Hi there how are you?"
	fmt.Println("Initial string is: " + sampString)

	// Returns true if phrase exists in string
	fmt.Println(strings.Contains(sampString, "are"))

	// Returns the index for the match
	fmt.Println(strings.Index(sampString, "ar"))

	// Returns the number of matches for the string
	fmt.Println(strings.Count(sampString, "e"))

	// Replaces the first letter with the second as many times as you define
	fmt.Println(strings.Replace(sampString, "e", "E", 2))

	// Return a list separating with the defined separator
	csvString := "1,2,3,4,5,6,7,8,9;10;11"
	//protect ; usage
	csvString = strings.Replace(csvString, ";", ",", -1)

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "d", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)

	// Returns a string using the values passed in separated with separator
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

	fmt.Println(listOfNums)

}

// FILE I/O
func fileOperation() {
	fmt.Println("###### File operations ######")

	// Create a file
	file, err := os.Create("samp.txt")

	// Output any errors
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("OK file created")
	}

	// Write a string to the file
	file.WriteString("Lets write some text to the file")

	// Close the file
	file.Close()

	// Try to open the file
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Convert into a string
	readString := string(stream)

	fmt.Println(readString)
}

// EXCEPTING INPUT
func getInput() {
	fmt.Println("###### working with user input ######")

	fmt.Println("What is your name? ")

	var name string

	fmt.Scan(&name)

	fmt.Println("Hello", name)

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
