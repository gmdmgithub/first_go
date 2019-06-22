package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	match, err := regexp.MatchString(`^p([a-z,_]+)ch`, "ap_some_text_ch")
	if err != nil {
		log.Printf("Probelm discovered: %v", err)
	}
	fmt.Println(match)

	matched, err := regexp.MatchString(`^a([a-z]+)b$`, "aa1xbb")
	fmt.Println(matched, err) // false

	a := regexp.MustCompile(`a`)
	fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
	fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
	fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
	fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
	fmt.Printf("%q\n", re.FindString("meat"))         // ""

	r, err := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.ReplaceAllString("a peach is a punch", "<fruit>"), err)
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
}
