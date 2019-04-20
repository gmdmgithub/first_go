package main

import (
	"log"
	"net/url"
	"os"
)

func main() {

	s := "postgres://uTest:uPass@myhost.com:5433/path?key=val&k2=val2#anchor"

	u, err := url.Parse(s)
	if err != nil {
		log.Printf("Problem with parsing %v\n", err)
		os.Exit(1)
	}
	log.Printf("Parsed url %+v", u)

	log.Printf("Parsed port %v", u.Port())
	log.Printf("Parsed scheme %v", u.Scheme)
	log.Printf("Parsed path %v", u.Path)
	log.Printf("Parsed query %+v", u.Query())
	log.Printf("Parsed query for key \"k\": %v", u.Query().Get("k"))
	log.Printf("Parsed user %v", u.User)
	log.Printf("Parsed user %v", u.User.Username())
	p, _ := u.User.Password()
	log.Printf("Parsed user %v", p)

	log.Printf("Path %v", u.Path)
	log.Printf("Anchor %v", u.Fragment)

}
