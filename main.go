package main

import (
	"fmt"
	"os"
)

func main() {
	paths := discoverPaths()
	args := os.Args[1:]

	suggestedPaths := fuzzyFindFile(args[0], paths)
	fmt.Printf(suggestedPaths[0])
}
