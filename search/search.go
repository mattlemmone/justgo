package search

import (
  "sort"
  "math"
  "fmt"
  "strings"
  "github.com/mattlemmone/popo/desktop"
)


type ApplicationRanking struct {
  Application *desktop.Application
  Score float64
}

type FileRanking struct {
  File *desktop.File
  Score float64
}

const (
  // lower number means longer directories will be penalized less
  directoryBonusNumberator = 0.7
)

func FuzzyFindApplication(target string, applications []*desktop.Application) []*desktop.Application {
  var rankings []ApplicationRanking
  var results []*desktop.Application
  
  loweredTarget := strings.ToLower(target)

  for _, app := range applications {
    name := strings.ToLower(app.Name)
    score := subsequenceSimilarity(loweredTarget, name)

    ranking := ApplicationRanking{
      Application: app,
      Score: score,
    }

    rankings = append(rankings, ranking)
  }

  sort.Slice(rankings, func(i, j int) bool{
    return rankings[i].Score > rankings[j].Score
  })

  for i := range rankings {
    results = append(results, rankings[i].Application)
  }

  return results

  // for i := range rankings {
  //   results = append(
  //     results,
  //     fmt.Sprintf("%s (%v)", rankings[i].Application.Name, rankings[i].Score),
  //   )
  // }

  // return results
}

func FuzzyFindFile(target string, files []*desktop.File) []string {
  var rankings []FileRanking
  var results []string
  
  loweredTarget := strings.ToLower(target)

  for _, file := range files {
    filepath := strings.ToLower(file.Path)

    if !subsequence(loweredTarget, filepath) {
      continue
    }

    dirs := strings.Split(filepath, "/")[1:]

    // add bonus to prefer short paths
    score := directoryBonusNumberator / math.Log(float64(len(dirs)))

    lastFile := dirs[len(dirs) - 1]
    score += subsequenceSimilarity(loweredTarget, lastFile)

    ranking := FileRanking{
      File: file,
      Score: score,
    }

    rankings = append(rankings, ranking)
  }

  sort.Slice(rankings, func(i, j int) bool{
    return rankings[i].Score > rankings[j].Score
  })

  for i := range rankings {
    results = append(
      results,
      fmt.Sprintf("%s (%v)", rankings[i].File.Path, rankings[i].Score),
    )
  }

  return results
}
