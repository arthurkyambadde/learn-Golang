package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//learning strings
	//strings are double quotes
	//variables in Go must be used or else we get an error

	var firstName string = "Arthur"
	// Go can infer for us a data type incase we dont explicitly assign one
	var secondName = "kyambadde"

	//we can also create empty strings in Go
	var freindname string

	fmt.Println("'First printing'", firstName, secondName, freindname)

	//In Go we can reassign our variables different figures

	firstName = "Nickson"
	secondName = "Daniel"
	freindname = "Jimmy"

	fmt.Println(" 'second printing' ", firstName, secondName, freindname)

	//when reassigning the datatype is supposed to be the same

	//Here is a Go short hand for creating a string

	freindJob := " web developer "

	fmt.Println("freind job", freindJob)

}
