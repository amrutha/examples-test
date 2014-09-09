// Package "bar" description goes here.

package bar

// Import some things.
import (
	"fmt"
)

// An unexported function will not do much this package outside.
func bar() {
	// But it's not of much use anyway.
	fmt.Println("Bar")
}
