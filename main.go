package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//learning strings
	//strings are double quotes
	//variables in Go must be used or else we get an error

	// var firstName string = "Arthur"
	// Go can infer for us a data type incase we dont explicitly assign one
	// var secondName = "kyambadde"

	//we can also create empty strings in Go
	// var freindname string

	// fmt.Println("'First printing'", firstName, secondName, freindname)

	//In Go we can reassign our variables different figures

	// firstName = "Nickson"
	// secondName = "Daniel"
	// freindname = "Jimmy"

	// fmt.Println(" 'second printing' ", firstName, secondName, freindname)

	//when reassigning the datatype is supposed to be the same

	//Here is a Go short hand for creating a string

	// freindJob := " web developer "

	// fmt.Println("freind job", freindJob)

	//Learn integers
	// if an integer is not assigned its automatically assigned 0

	var age1 int = 20

	var age2 = 10

	var age3 int

	age4 := 30

	fmt.Println(age1, age2, age3, age4)

	//refer to https://pkg.go.dev/builtin#int for types of integers
	//The uint only declares positive numbers

	var numOne unit = 25

	//bit and memory
	//we have int8, int16, int32 and int 64
	//all can allow a given range of numbers in inreasing order and the same applies
	//to uint we have uint8, uint16, uint32, uint64

	var number2 uint8 = 224

	//Floats
	//Floats allow us to use decimal places on numbers
	//we have float32 and float 64
	//https://go.dev/ref/spec#Numeric_types

	fmt.Println(numOne, number2)

}
