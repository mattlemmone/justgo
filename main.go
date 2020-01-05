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
  // apps := environment.DesktopApplications()
  userFiles := environment.UserFiles()
  // files := append(apps, userFiles...)

  reader := bufio.NewReader(os.Stdin)
  
  var seed []string
  for i := range userFiles {
   seed = append(seed, userFiles[i].Path) 
  }

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    result := search.FuzzyFindFile(text[:len(text)-1], seed)

    for i := range result[:20] {
      path := result[i]
      println(path)
    }
  }
}

