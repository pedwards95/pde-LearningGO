package main

import "fmt"

func main() {
	var you string = "Paul"
	fmt.Println("Hey", you+", how are you?")

	const earthsGravity = 9.80665
	fmt.Println(earthsGravity)

	var kilometersToMars int32 = 62100000
	fmt.Println(kilometersToMars)

	daysOnVacation := 7
	var hoursInDay = 24
	const DAYS int = 3
	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation. Weren't you supposed to go back to work after", DAYS, "days?")

	var basketTotal float64
	spinachPrice := 1.50
	basketTotal += spinachPrice
	fmt.Println(basketTotal)

	var part1, part2 string
	part1 = "To be..."
	part2 = "Not to be..."
	fmt.Println(part1, part2)

	quote, fact := "Bears, Beets, Battlestar Galactica", true
	fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
	fmt.Println(fact)  // Prints: true
}
