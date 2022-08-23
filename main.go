package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	// fmt.Println("hello world")
// 	//learning strings
// 	//strings are double quotes
// 	//variables in Go must be used or else we get an error

// 	// var firstName string = "Arthur"
// 	// Go can infer for us a data type incase we dont explicitly assign one
// 	// var secondName = "kyambadde"

// 	//we can also create empty strings in Go
// 	// var freindname string

// 	// fmt.Println("'First printing'", firstName, secondName, freindname)

// 	//In Go we can reassign our variables different figures

// 	// firstName = "Nickson"
// 	// secondName = "Daniel"
// 	// freindname = "Jimmy"

// 	// fmt.Println(" 'second printing' ", firstName, secondName, freindname)

// 	//when reassigning the datatype is supposed to be the same

// 	//Here is a Go short hand for creating a string

// 	// freindJob := " web developer "

// 	// fmt.Println("freind job", freindJob)

// 	//Learn integers
// 	// if an integer is not assigned its automatically assigned 0

// 	// var age1 int = 20

// 	// var age2 = 10

// 	// var age3 int

// 	// age4 := 30

// 	// fmt.Println(age1, age2, age3, age4)

// 	//refer to https://pkg.go.dev/builtin#int for types of integers
// 	//The uint only declares positive numbers

// 	// var numOne unit = 25

// 	//bit and memory
// 	//we have int8, int16, int32 and int 64
// 	//all can allow a given range of numbers in inreasing order and the same applies
// 	//to uint we have uint8, uint16, uint32, uint64

// 	// var number2 uint8 = 224

// 	//Floats
// 	//Floats allow us to use decimal places on numbers
// 	//we have float32 and float 64
// 	//https://go.dev/ref/spec#Numeric_types

// 	// fmt.Println(numOne, number2)

// 	// ----------------------------------------------------------------------
// 	// -----------------------------------------------------------------------
// 	// ----------------------------------------------------------------------------

// 	//learning fmt package functions

// 	//1) Print function

// 	//it prints strings on the same line no matter
// 	//how you print them in your function
// 	//In order to print on a new line
// 	//we add "\n"
// 	// fmt.Print("Hello, \n")
// 	// fmt.Print("world!")

// 	//2) Println function

// 	//it prints strings on different lines
// 	// fmt.Println("My names is Arthur")
// 	// fmt.Println("I am a Front-end dev")

// 	// name := "Arthur"
// 	// age := 45

// 	//Formatted strings
// 	//Printf(formatted strings) %_ is a format specifier
// 	//it prints strings on the same line

// 	//%v prints variables in a default manner
// 	// fmt.Printf("My name is %v and am aged %v\n", name, age)

// 	//%q add doublue quotes to strings variables and turns integers variables to #
// 	// fmt.Printf("My name is %q and am aged %q\n", name, age)

// 	//%T gets the type of a variable
// 	// fmt.Printf("My name is %T and am aged %T\n", name, age)

// 	//%f out puts float numbers and we can specify decimal places and it rounds off numbers

// 	// fmt.Printf("You have scored %0.2f\n", 45.78563)

// 	//Sprint (save formatted string)

// 	// var str = fmt.Sprintf("my name is %v", name)
// 	// fmt.Println(str)

// 	//'###################################################################
// 	//######################################################################
// 	//#####################################################################

// 	//ARRAYS IN GO

// 	//we declare an array as a variable
// 	//specify its length in sqaure brackets and specify the data type
// 	//it has a fixed length

// 	// var freinds [5]string = [5]string{
// 	// 	"jimmy", "daniel", "denis", "junior", "jordan",
// 	// }

// 	//Go can also infer for us a data type of an array

// 	// var newFriends = [5]string{
// 	// 	"jimmy", "daniel", "denis", "junior", "jordan",
// 	// }

// 	// fmt.Printf("%T", newFriends)

// 	// fmt.Println(newFriends)

// 	//length of an array
// 	// fmt.Println(len(newFriends))

// 	// ======================================================
// 	//=========================================================
// 	//=======================================================

// 	//Slices
// 	//they use arrays under the hood

// 	// var scores = []int{20, 30, 40, 50}
// 	// scores[2] = 45

// 	// //add/append an item to a slice
// 	// //the append slice method returns a new slice with an added value
// 	// //the added item is added to the end of the slice

// 	// scores = append(scores, 63)

// 	//lenght of a slice

// 	// lengthScores := len(scores)

// 	//slice ranges
// 	//these return a range of values from a slice
// 	//it returns the value at first specified index to the value
// 	//before the second specified index but not the value at
// 	//the second specified index

// 	// rangeOne := scores[1:3]
// 	// rangetwo := scores[0:7]
// 	// rangethree := scores[2:] //return all values from position 2 to the end of the slice
// 	// rangefour := scores[:3]  // retun values from start to position 3 but not including position 3

// 	// fmt.Println("range one :", rangeOne)
// 	// fmt.Println("range two :", rangetwo)
// 	// fmt.Println("range three :", rangethree)
// 	// fmt.Println("range four :", rangefour)

// 	// // scores = append(scores, rangeOne...)
// 	// // fmt.Println("scores : ", scores)

// 	// //Starndard library

// 	// //Strings methods
// 	// //They return a manipulated copy but they donot mutate the original variable

// 	// // greeting := "good morning everyone"

// 	// //Contains
// 	// // fmt.Println(strings.Contains(greeting, "one"))
// 	// // fmt.Println(strings.Contains(greeting, "why"))

// 	// // //ReplaceAll
// 	// // fmt.Println(strings.ReplaceAll(greeting, "everyone", "Nickson"))

// 	// // //ToUpper
// 	// // fmt.Println(strings.ToUpper(greeting))

// 	// //Index
// 	// // fmt.Println(strings.Index(greeting, "o"))

// 	// //Split
// 	// //It returns a slice of the split string
// 	// // fmt.Println(strings.Split(greeting, " "))

// 	// ages := []int{
// 	// 	12, 54, 23, 21, 54, 82, 1, 5,
// 	// }

// 	// //Sort
// 	// //it arrages integers in increasing order
// 	// //it arranges strings in alphabetical order
// 	// sort.Ints(ages)
// 	// fmt.Println(ages)

// 	// var names = []string{
// 	// 	"jimmy", "daniel", "denis", "junior", "jordan",
// 	// }
// 	// sort.Strings(names)
// 	// fmt.Println(names)

// 	// fmt.Println(sort.SearchStrings(names, "jimmy"))

// 	// '''''''''''''''''''''''''''''''''''''''''''''''''''
// 	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// 	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

// 	//LOOPS IN GO

// 	// age := 0

// 	// for age < 5 {
// 	// 	fmt.Println("The age is : ", age)
// 	// 	age++
// 	// }

// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)

// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println("The value of i is : ", i)
// 	// }

// 	//Looping through a slice

// 	// names := []string{
// 	// 	"jimmy", "daniel", "denis", "junior", "jordan",
// 	// }

// 	// // for i := 0; i < len(names); i++ {
// 	// // 	fmt.Println("The value at position  i is : ", names[i])
// 	// // }

// 	// for index, value := range names {
// 	// 	fmt.Printf("The name at position %v is %v \n", index, value)
// 	// }

// 	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// 	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// 	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// 	//Functions in Go

// }

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func saybye(n string) {
// 	fmt.Printf("Good bye %v \n", n)
// }

// func sayGreetingAgain(n string) {
// 	fmt.Printf("Good evening %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, value := range n {
// 		f(value)
// 	}
// }

// //functions that return a value

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {

// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string

// 	for _, value := range names {
// 		initials = append(initials, value[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"

// }

// func main() {

// 	// names := []string{
// 	// 	"jimmy", "daniel", "denis", "junior", "jordan",
// 	// }

// 	// sayGreeting("daniel")
// 	// sayGreetingAgain(("enock"))

// 	// cycleNames(names, saybye)

// 	// areaOne := circleArea(12)
// 	// areaTwo := circleArea(12.5)

// 	// fmt.Println(areaOne, areaTwo)
// 	// fmt.Printf("circle one area %0.2f", areaOne)

// 	firstInitiaal1, secondInitial1 := getInitials(("Kyambadde Arthur"))
// 	firstInitiaal2, secondInitial2 := getInitials(("Naluyima sandra Bianca"))
// 	firstInitiaal3, secondInitial3 := getInitials(("when"))

// 	fmt.Println(firstInitiaal1, secondInitial1)
// 	fmt.Println(firstInitiaal2, secondInitial2)
// 	fmt.Println(firstInitiaal3, secondInitial3)

// }

//Go Scope
//values of one file can be used in another but the files are supposed to both
//run for the action to work

// func main() {

// 	sayHello("Arthur")

// 	for _, v := range points {
// 		fmt.Println(v)
// 	}

// }

// func main() {
// 	//Learning maps
// 	//its more of an object in javascript but
// 	//a different syntax

// 	// menu := map[string]float64{
// 	// 	"katogo": 2.99,
// 	// 	"rice":   10.65,
// 	// 	"irish":  5.22,
// 	// }

// 	phoneBook := map[int]string{
// 		77275675:  "Nickson",
// 		77436372:  "daniel",
// 		703764574: "jimmy",
// 	}

// 	phoneBook[703764574] = "jordan"

// 	// fmt.Println(menu)
// // 	// fmt.Println(menu["katogo"])

// // 	//Looping through maps

// // 	// for key, value := range menu {
// // 	// 	fmt.Println(
// // 	// 		"Our menu today contains -",
// // 	// 		key, ":", value,
// // 	// 	)
// // 	// }

// // 	fmt.Println(phoneBook)
// // 	fmt.Println(phoneBook[703764574])
// // }

// //Pointers
// //we use an ampasand(&) sign to point at a memory address
// //we place it before the variable name
// //We use the astrix sign(*) to read the value at the memory address

// func updateName(n *string) {
// 	*n = "Jimmy"

// }

// func updateTeam(n *string) {
// 	*n = "Arsenal"
// }

// func main() {
// 	name := "Arthur"
// 	m := &name

// 	team := "barca"

// 	m2 := &team

// 	updateName(m)
// 	updateTeam(m2)

// 	fmt.Println("memory location for names is : ", &name, m)
// 	fmt.Println(team)

// 	// fmt.Println(name)

// 	fmt.Println("value at memory address is : ", *m)
// }

// func main() {
// 	myBill := newBill("Arthur's Bill")

// 	myBill.updateTip(15)

// 	myBill.addItem("Katogo", 23.5)
// 	myBill.addItem("molokoni", 10.5)
// 	myBill.addItem("kyankya", 12.8)
// 	myBill.addItem("mayuuni", 7.5)
// 	myBill.addItem("mukene", 9.5)

// 	fmt.Println(myBill.format())

// }

func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("create a new bill: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("create a new bill: ", reader)

	b := newBill(name)

	fmt.Println("created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a--add an item, s--save the bill,t--add a tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Whats the item name: ", reader)
		price, _ := getInput("Whats the item price: ", reader)

		prc, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The bill must be a number ")
			promptOptions(b)
		}

		b.addItem(name, prc)

		fmt.Print("Item added to the bill", name, price)

	case "t":
		tip, _ := getInput("whats the tip amount ($):", reader)

		tp, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The bill must be a number ")
			promptOptions(b)
		}

		b.updateTip(tp)

		fmt.Print("tip has been added ", tip)

	case "s":
		b.saveBill()

		fmt.Print("you have saved the file ", b.name)

	default:
		fmt.Println("Not a valid option ........")
		promptOptions(b)

	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
