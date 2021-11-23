package main

import "fmt"

func main() {
	fmt.Println("Start")
	x := "My very first address"
	fmt.Println(&x)

	star := "Polaris"
	starAddress := &star
	fmt.Println("The address of star is", starAddress)

	lyrics := "Moments so dear"
	pointerForStr := &lyrics
	*pointerForStr = "Journeys to plan"
	fmt.Println(lyrics) // Prints: Journeys to plan

	greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}

func brainwash(saying *string) {
	*saying = "Beep Boop."
}
