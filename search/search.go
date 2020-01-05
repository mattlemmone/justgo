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

  for i := range filepaths {
    filepath := strings.ToLower(filepaths[i])
    dirs := strings.Split(filepath, "/")[1:]

    fileScore := 0.0

    if subsequence(loweredTarget, filepath) {
      fileScore += 1.0

      // short path bonus
      fileScore += float64(2.0) / float64(len(dirs))
    }

    lastFile := strings.ToLower(dirs[len(dirs) - 1])

    fileScore += subsequenceSimilarity(loweredTarget, lastFile)

    ranking := Ranking{
      Value: filepaths[i],
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
