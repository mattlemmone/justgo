package search

import (
  "sort"
  "strings"
)


type Ranking struct {
  Value string
  Score float64
}


func FuzzyFind(target string, choices ...string) []string {
  var rankings []Ranking
  var rankedChoices []string

  loweredTarget := strings.ToLower(target)

  for i := range choices {
    score := levenshteinDistance(
      loweredTarget,
      strings.ToLower(choices[i]),
    )

    ranking := Ranking{Value: choices[i], Score: score}
    rankings = append(rankings, ranking)
  }

  sort.Slice(rankings, func(i, j int) bool{
    return rankings[i].Score > rankings[j].Score
  })

  for i := range rankings {
    rankedChoices = append(rankedChoices, rankings[i].Value)
  }

  return rankedChoices
}

func FuzzyFindFile(path string, choices ...string) []string {
  // todo: split paths and add to choices
  return []string{}
}
