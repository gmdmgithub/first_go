package main

import (
	"fmt"
	"sort"
	"strings"
)

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
	csvString = strings.Replace(csvString, ";", ",", -1) //-1 replace all - from version 1.12 it is promise ReplaceAll func

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "d", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)

	// Returns a string using the values passed in separated with separator
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

	fmt.Println(listOfNums)

}
