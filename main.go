package main

import (
	"fmt"
	"os"
	// "os/exec"
	"bufio"
	"github.com/mattlemmone/popo/search"
	// "github.com/mattlemmone/popo/gui"
	"github.com/mattlemmone/popo/desktop"
)

func main() {
	environment := desktop.Linux{}

	println("Scanning...")
	apps := environment.DesktopApplications()
	// userFiles := environment.UserFiles()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// results := search.FuzzyFindFile(text[:len(text)-1], userFiles)
		results := search.FuzzyFindApplication(text[:len(text)-1], apps)

		for i := range results[:5] {
			fmt.Println(results[i].Name)
		}

		topResult := results[0]
		environment.LaunchApplication(topResult)
		break
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}
