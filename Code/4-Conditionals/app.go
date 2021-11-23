package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	heistReady := false

	if heistReady {
		fmt.Println("Let's go!")
	} else {
		fmt.Println("Act normal.")
	}
	fmt.Println("password1" == "password1") // Evaluates to true
	fmt.Println("password1" == "badpass")   // Evaluates to false

	rightTime := true
	rightPlace := true
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}

	name := "H. J. Simp"

	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}

	amountLeft := rand.Intn(10000)
	fmt.Println("amountLeft is: ", amountLeft)
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}

	rand.Seed(time.Now().UnixNano()) //THIS MAKES IT NOT THE SAME NUMBER EVERY TIME. UNIX NANO IS SECONDS FROM JAN 1 1970,
	// MEANING THIS NUMBER WILL ALWAYS BE GOING UP, AND THUS ALWAYS UNIQUE
	// Go uses a default seed value of 1 which can lead to predictable numbers being generated. We can generate random numbers by using different seed values.
	amountLeft = rand.Intn(10000)
	fmt.Println("amountLeft is: ", amountLeft)
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}
