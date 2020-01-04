package main

import (
  "fmt"
  "github.com/mattlemmone/popo/search"
)


func main(){
  target := "test"
  choices := []string{"hey", "test"}
  results := search.FuzzyFind(target, choices...)

  fmt.Printf("results: %s", results)
}

