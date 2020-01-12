package main


func main() {
  paths := discoverPaths()
  suggestedPaths := fuzzyFindFile("main", paths)
  println(suggestedPaths)
}
