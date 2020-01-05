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
  userFiles := environment.UserFiles()
  files := append(apps, userFiles...)

  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    result := search.FuzzyFindFile(text[:len(text)-1], files)

    for i := range result[:10] {
      file := result[i]
      println(file.Path)
    }
  }
}

