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


func main(){
  environment := desktop.Linux{}

  println("Scanning...")
  apps := environment.DesktopApplications()

  for i := range apps {
    println(apps[i].Name)
  }
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

    fmt.Printf("> %s\n", topResult.Exec)

    // parsed := append([]string{"-c"}, topResult.Exec)
    // cmd := exec.Command("bash", parsed...)
    // out, err := cmd.Output()
    // println(string(out))
    // if err != nil {
    //   panic(err)
    // }

    break
    // go cmd.Start()

    // for i := range result[:min(10,len(result))] {
    //   path := result[i]
    //   println(path)
    // }
  }
}


func min(a int, b int) int{
  if a > b {
    return b
  }

  return a
}
