package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hi there, here we have in memory database ############")
	defer fmt.Println("#################  Bey!!, end of in memory db")
	db := NewStore()
	defer db.Clear()

	db.Insert("Alex", "test")
	db.Insert("Max", "Hi there")
	db.Insert("Alex", "best?")
	printDB(db)
	db.Update("Alex", "Now update!")
	db.Insert("John", "John is rolling!")
	printDB(db)
	db.Delete("Alex")
	printDB(db)

}

// KeyVal - struct to keep inmemory DB
type KeyVal struct {
	kv map[string]string
}

// NewStore - create new DB store
func NewStore() *KeyVal {
	return &KeyVal{
		kv: map[string]string{},
	}
}

// Insert - insert a db
func (db *KeyVal) Insert(k, v string) {
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
func (db *KeyVal) Update(k, v string) {
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
