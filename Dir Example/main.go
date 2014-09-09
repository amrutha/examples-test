// Package "main" description.

package main

import (
	"fmt"

	// Import our package.
	"utils"
)

func main() {
	// Use function "Foo" into package "utils".
	str, err := utils.Foo()
	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}
