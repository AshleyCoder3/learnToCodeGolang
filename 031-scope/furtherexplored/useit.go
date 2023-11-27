package furtherexplored

import "fmt"

// this is also "package block" scope
// this is exported because the identifier "Fascinating" is capitalized

func Fascinating() {
	fmt.Println("Planet speed:", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet spinning speed:", planetRotationSpeed)
}
