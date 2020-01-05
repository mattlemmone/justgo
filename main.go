package main

import (
  "fmt"
  "os"
  "bufio"
  "github.com/mattlemmone/popo/search"
  // "github.com/mattlemmone/popo/gui"
  "github.com/mattlemmone/popo/desktop"
)


func main(){
  environment := desktop.Linux{}

  println("Scanning...")
  apps := environment.DesktopApplications()
  // userFiles := environment.UserFiles()

  reader := bufio.NewReader(os.Stdin)
  
  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // result := search.FuzzyFindFile(text[:len(text)-1], userFiles)
    result := search.FuzzyFindApplication(text[:len(text)-1], apps)

    for i := range result[:20] {
      path := result[i]
      println(path)
    }
  }
}

