package main

import "fmt"

func main() {

	// Declaration
	var favBook = "Harry Potta"

	// Reassignment
	favBook = "Atomic habits"

	fmt.Println(favBook)

	// Declaration with type
	var otherFavBook string

	otherFavBook = "Bad Blood"

	fmt.Println(otherFavBook)

	// Block declaration
	var (
		favNumber=1
		favChocolate="bounty"
	)
	fmt.Println(favNumber, favChocolate)

	// Declaration with colon
	favAnimal := "cat"
	fmt.Println(favAnimal)

	// Constants 
	const myName = "Filip"
	// cant do that myName = "Jan"
}