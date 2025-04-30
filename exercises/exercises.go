package exercises

import "fmt"

// exerciseRunner defines the type for functions that run exercise tests.
type exerciseRunner func()

// exerciseMap stores the mapping between exercise identifiers and their runner functions.
var exerciseMap = map[string]exerciseRunner{
	"20":  RunValidParenthesesTests,
	"121": RunValidMaxProfitTest,
	// Add other exercises here
}

// RunExercise looks up the exercise by its identifier and runs it if found.
func RunExercise(identifier string) bool {
	runner, found := exerciseMap[identifier]
	if !found {
		fmt.Printf("Exercise '%s' not found.\n", identifier)
		fmt.Println("Available exercises: validParenthesis (or 20)") // TODO: Make this dynamic
		return false
	}
	runner()
	return true
}
