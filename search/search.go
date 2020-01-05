package search

import (
  "sort"
  "fmt"
  "strings"
)


type Ranking struct {
  Value string
  Score float64
}


func FuzzyFindFile(target string, filepaths []string) []string {
  var rankings []Ranking
  var results []string
  
  loweredTarget := strings.ToLower(target)
  // cache := map[string]float64{}

  for i := range filepaths {
    filepath := strings.ToLower(filepaths[i])
    dirs := strings.Split(filepath, "/")[1:]

    fileScore := 0.0

    // for j := range dirs {
    //   dir := dirs[j]

    //   // if subsequence(loweredTarget, dir) {
    //   //   fmt.Printf("yay %s\n", dir)
    //   //   fileScore += .1
    //   // }

    //   dirScore, exists := cache[dir]

    //   if !exists {
    //     if len(target) > 2 {
    //       dirScore = levenshteinDistance(
    //         loweredTarget,
    //         strings.ToLower(dir),
    //       )
    //     } else {
    //       dirScore = cosine(
    //         loweredTarget,
    //         strings.ToLower(dir),
    //       )
    //     }

    //     cache[dir] = dirScore
    //   }

    //   fileScore += dirScore 
    // }

    // fileScore /= float64(len(dirs))

    lastFile := dirs[len(dirs) - 1]

    if lastFile  == loweredTarget {
      fileScore += 1
      fmt.Printf("%+v", fileScore)
    }

    // if subsequence(loweredTarget, lastFile){
    //   fileScore += .2
    //   fmt.Printf("%+v", fileScore)
    // }

    // file.Path += fmt.Sprintf(" %s", fileScore)
    ranking := Ranking{
      Value: filepath,
      Score: fileScore,
    }

    rankings = append(rankings, ranking)
  }

  sort.Slice(rankings, func(i, j int) bool{
    return rankings[i].Score > rankings[j].Score
  })

  for i := range rankings {
    results = append(
      results,
      fmt.Sprintf("%s (%v)", rankings[i].Value, rankings[i].Score),
    )
  }

  return results
}
