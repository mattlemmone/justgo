package search

import (
  "sort"
  "fmt"
  "strings"
  "github.com/mattlemmone/popo/desktop"
)


type Ranking struct {
  Value string
  Score float64
}

type FileRanking struct {
  File *desktop.File
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

func FuzzyFindFile(target string, files []*desktop.File) []*desktop.File {
  var rankings []FileRanking
  var rankedChoices []*desktop.File
  
  loweredTarget := strings.ToLower(target)
  cache := map[string]float64{}

  for i := range files {
    file := files[i]
    dirs := strings.Split(file.Path, "/")[1:]
    fileScore := 0.0

    for j := range dirs {
      dir := dirs[j]
      dirScore, exists := cache[dir]
      
      if !exists {
        dirScore = levenshteinDistance(
          loweredTarget,
          strings.ToLower(dir),
        )

        fmt.Printf("score %s %+v\n", dir, dirScore)
        cache[dir] = dirScore
      }
           
      fileScore += dirScore 
    }

    fileScore += float64(1.0 / len(dirs))

    if dirs[len(dirs) - 1] == loweredTarget {
      fmt.Println("BITCH")
      fileScore += 3
      fmt.Printf("%+v", fileScore)
    }

    file.Path += fmt.Sprintf(" %s", fileScore)
    ranking := FileRanking{
      File: file,
      Score: fileScore,
    }

    rankings = append(rankings, ranking)
  }

  sort.Slice(rankings, func(i, j int) bool{
    return rankings[i].Score > rankings[j].Score
  })

  for i := range rankings {
    rankedChoices = append(rankedChoices, rankings[i].File)
  }

  return rankedChoices
}
