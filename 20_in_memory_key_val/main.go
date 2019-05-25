package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
	age  int
}

func main() {
	fmt.Println("############ Hi there, here we have in memory database")
	defer fmt.Println("#################  Bey!!, end of in memory db")
	db := NewStore()
	defer db.Clear() //not necessary but nice

	db.Insert("Alex", "test")
	db.Insert("Max", "Hi there")
	db.Insert("Alex", "best?")

	printDB(db)

	db.Update("Alex", "Now update!")
	db.Insert("John", "John is rolling!")
	db.Insert("Morgan", 20)
	db.Insert("Sara", 18.3)
	db.Insert("Megan", user{"Frank", 12})

	printDB(db)

	db.Delete("Alex")

	printDB(db)

}

// KeyVal - struct to keep in-memory DB
type KeyVal struct {
	kv map[string]interface{}
}

// NewStore - create new DB store
func NewStore() *KeyVal {
	return &KeyVal{
		kv: make(map[string]interface{}),
	}
}

// Insert - insert a db
func (db *KeyVal) Insert(k string, v interface{}) {
	if _, ok := db.kv[k]; ok {
		log.Println("Key exists, no insert", k)
		return
	}
	db.kv[k] = v
}

// Clear - clear db
func (db *KeyVal) Clear() {
	log.Println("clear db")
	db = nil
}

func printDB(db *KeyVal) {

	if db == nil {
		fmt.Println("No data")
		return
	}

	fmt.Println("############ PRINTING DB ##############")

	for k, v := range db.kv {
		fmt.Println("DB data", k, v)
	}
}

// Update - update keyVal db
func (db *KeyVal) Update(k string, v interface{}) {
	if _, ok := db.kv[k]; !ok {
		log.Println("no key to update", k)
		return
	}
	db.kv[k] = v
}

// Delete - delete element if exist
func (db *KeyVal) Delete(k string) {
	if _, ok := db.kv[k]; ok {
		delete(db.kv, k)
	}
}
