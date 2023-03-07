package main

import "fmt"

func main(){
	// Arithmetiic operator
	// + - / * %
	var num1 = 5
	var num2 = 2

	var sum = num1 + num2
	var difference = num1 - num2
	var quotient = num1 / num2
	var product = num1 * num2
	var remainder = num1 % num2
	fmt.Println(sum, difference, quotient, product, remainder)

	// Relational operator
	// > < == != || &&
	var result = num1 == 1 || num2 == 2

	fmt.Println(result)

	// Logical operator
	const name = "Laith"
	var age = 25

	var inviteToParty = name == "Laith" && age > 21

	fmt.Println(inviteToParty)
}