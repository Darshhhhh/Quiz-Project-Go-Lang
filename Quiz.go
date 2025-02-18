package main

import "fmt"

func main() {
	var UserName string
	fmt.Printf("Enter Your Name: ")
	fmt.Scan(&UserName)
	fmt.Printf("Hey! %v,Ready to Play?\n", UserName)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	// fmt.Println(age >= 10) //gives true false

	if age >= 10 {
		fmt.Println("Yoo! You can play this game!")
	} else {
		fmt.Println("Opps! You cant Play this Game!")
		return
	}
	fmt.Printf("What is better, the RTX 3080 or RTX 3090?")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		fmt.Printf("You Won")
	} else {
		fmt.Println("Incorrect!")
	}
}
