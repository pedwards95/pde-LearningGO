package main

import "fmt"

func main() {
	eatTacos()
	myAge := 25

	myMartianAge := computeMarsYears(myAge)
	fmt.Println(myMartianAge)

	var likes, shares int

	// Update this line so the results of the function
	// get stored in "likes" and "shares"
	likes, shares = getLikesAndShares(4)

	if likes > 5 {
		fmt.Println("Woohoo! We got some likes.")
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}

	queryDatabase("SELECT * FROM coolTable;")
}

func eatTacos() {
	fmt.Println("Yum!")
}

func computeMarsYears(earthYears int) int {

	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func getLikesAndShares(postId int) (int, int) {
	var likesForPost, sharesForPost int
	switch postId {
	case 1:
		likesForPost = 5
		sharesForPost = 7
	case 2:
		likesForPost = 3
		sharesForPost = 11
	case 3:
		likesForPost = 22
		sharesForPost = 1
	case 4:
		likesForPost = 7
		sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
	// Add in a return for likesForPost and sharesForPost
	return likesForPost, sharesForPost

}

func queryDatabase(query string) string {
	var result string
	connectDatabase()
	// Add deferred call to disconnectDatabase() here
	defer disconnectDatabase()

	if query == "SELECT * FROM coolTable;" {
		result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}
