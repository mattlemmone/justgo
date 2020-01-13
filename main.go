package main

import (
  "fmt"
  "os"
	"os/exec"
)

const maxSearchResults = 5

func main() {
  paths := discoverPaths()
  // fmt.Printf("discovered: %+v\n", paths)
  args := os.Args[1:]
  suggestedPaths := fuzzyFindFile(args[0], paths)
  numResults := min(maxSearchResults, len(suggestedPaths))
  fmt.Printf("suggestedPaths: %+v\n", suggestedPaths[:numResults])
  echo(suggestedPaths[0])
}

func echo(s string){
    sh := os.Getenv("SHELL") 

		output := fmt.Sprintf("echo %s", s)

    cmd := exec.Command(sh, "-c ", output)
    cmd.Stdout = os.Stdout
    cmd.Run()
}
