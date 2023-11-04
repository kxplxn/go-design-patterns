package main

import "fmt"

// Iterate using a callback function
func main() {
	// use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// use IterateBooks to iterate via anonymous function
	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})

	// TODO: create a BookIterator
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
