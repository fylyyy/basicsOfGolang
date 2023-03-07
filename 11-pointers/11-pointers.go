package main

import "fmt"

func main(){
	var a = 1

	var b = a

	b++

	var c = 1

	var d *int

	// & specifying pointer to specific variable
	d = &c

	// use star if you want to access a pointer, fmt.Println(d) prints out the pointer value
	*d += 1

	fmt.Println(a, c, d)
}