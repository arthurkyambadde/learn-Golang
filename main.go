package main

import "fmt"

func main() {
	// fmt.Println("hello world")
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

	// var age1 int = 20

	// var age2 = 10

	// var age3 int

	// age4 := 30

	// fmt.Println(age1, age2, age3, age4)

	//refer to https://pkg.go.dev/builtin#int for types of integers
	//The uint only declares positive numbers

	// var numOne unit = 25

	//bit and memory
	//we have int8, int16, int32 and int 64
	//all can allow a given range of numbers in inreasing order and the same applies
	//to uint we have uint8, uint16, uint32, uint64

	// var number2 uint8 = 224

	//Floats
	//Floats allow us to use decimal places on numbers
	//we have float32 and float 64
	//https://go.dev/ref/spec#Numeric_types

	// fmt.Println(numOne, number2)

	// ----------------------------------------------------------------------
	// -----------------------------------------------------------------------
	// ----------------------------------------------------------------------------

	//learning fmt package functions

	//1) Print function

	//it prints strings on the same line no matter
	//how you print them in your function
	//In order to print on a new line
	//we add "\n"
	// fmt.Print("Hello, \n")
	// fmt.Print("world!")

	//2) Println function

	//it prints strings on different lines
	// fmt.Println("My names is Arthur")
	// fmt.Println("I am a Front-end dev")

	name := "Arthur"
	age := 45

	//Formatted strings
	//Printf(formatted strings) %_ is a format specifier
	//it prints strings on the same line

	//%v prints variables in a default manner
	fmt.Printf("My name is %v and am aged %v\n", name, age)

	//%q add doublue quotes to strings variables and turns integers variables to #
	fmt.Printf("My name is %q and am aged %q\n", name, age)

	//%T gets the type of a variable
	fmt.Printf("My name is %T and am aged %T\n", name, age)

	//%f out puts float numbers and we can specify decimal places and it rounds off numbers

	fmt.Printf("You have scored %0.2f\n", 45.78563)

	//Sprint (save formatted string)

	var str = fmt.Sprintf("my name is %v", name)
	fmt.Println(str)

}
