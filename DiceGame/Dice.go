package main

import (
	"fmt"
	"math/rand"
)

// func getRandomN(min int, max int) int {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	rng := rand.New(source)

//		return rand.Intn(max-min) + min
//	}
func main() {
	bot_number := rand.Intn(5) + 1
	player_number := rand.Intn(5) + 1
	result := ""

	if bot_number > player_number {
		result = "Bot won"
	} else if player_number > bot_number {
		result = "Player won"
	} else {
		result = "It's tie!"
	}

	fmt.Println("Player's number:", player_number)
	fmt.Println("Bot's number:", bot_number)
	fmt.Println("Result:", result)

}
