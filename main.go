package main

import (
	"leetcode/exercises"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide the exercise name or number as an argument.")
	// 	// TODO: Potentially list available exercises dynamically from the map
	// 	return
	// }

	// exerciseIdentifier := os.Args[1]

	// Call the central exercise runner
	if !exercises.RunExercise("125") {
		// Optional: Handle the case where the exercise was not found,
		// though RunExercise already prints a message.
	}
}
