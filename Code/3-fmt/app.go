package main

import "fmt"

func main() {
	votingAge := 18
	fmt.Printf("You must be %d years old to vote.", votingAge)
	fmt.Println("\n***") // Added for spacing

	specialNum := 42
	fmt.Printf("This value's type is %T.", specialNum)
	fmt.Println("\n***") // Added for spacing

	gpa := 3.8
	fmt.Printf("You're averaging: %.2f.", gpa)
	fmt.Println("\n***") // Added for spacing

	floatExample := 1.75
	fmt.Printf("Working with a %T", floatExample)
	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years..", yearsOfExp, reqYearsExp)
	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	fmt.Printf("Each share of Gopher feed is $%.2f", stockPrice)
	fmt.Println("\n***") // Added for spacing
	fmt.Printf("Each share of Gopher feed is $%.1f", stockPrice)
	fmt.Println("\n***") // Added for spacing
	fmt.Printf("Each share of Gopher feed is $%.3f", stockPrice)
	fmt.Println("\n***") // Added for spacing

	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	fmt.Println(teacherSays) // Prints: You scored a 100 on the test! Great job!

	template := "I wish I had a %v."
	pet := "puppy"
	var wish string
	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)

	fmt.Println("How are you doing?")

	var response1 string
	var response2 string
	fmt.Scan(&response1)
	fmt.Scan(&response2)
	fmt.Printf("I'm %v %v", response1, response2)
}
