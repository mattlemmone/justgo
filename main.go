package main

import (
	"fmt"
	"os"
	// "os/exec"
)

func main() {
	paths := discoverPaths()
	// fmt.Printf("discovered: %+v\n", paths)
	args := os.Args[1:]
	suggestedPaths := fuzzyFindFile(args[0], paths)
	fmt.Printf(suggestedPaths[0])
}

// func echo(s string) {
// 	sh := os.Getenv("SHELL")

// 	output := fmt.Sprintf("echo %s", s)

// 	cmd := exec.Command(sh, "-c ", output)
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }
