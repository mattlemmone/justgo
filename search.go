package main

import (
	// "fmt"
	"math"
	"sort"
	"strings"
)

// FileRanking is used to sort different files based on their score assigned in fuzzyFindFile.
type FileRanking struct {
	Path  string
	Score float64
}

const (
	// lower number means longer directories will be penalized less
	directoryBonusNumberator = 0.7
)

func fuzzyFindFile(target string, paths []string) []string {
	var rankings []FileRanking
	var results []string

	loweredTarget := strings.ToLower(target)

	for _, path := range paths {
		loweredPath := strings.ToLower(path)

		if !Subsequence(loweredTarget, loweredPath) {
			continue
		}

		// todo: dont split, substr
		dirs := strings.Split(loweredPath, "/")[1:]

		// add bonus to prefer short paths
		score := directoryBonusNumberator / math.Log(float64(len(dirs)))

		lastFile := dirs[len(dirs)-1]
		score += SubsequenceSimilarity(loweredTarget, lastFile)

		ranking := FileRanking{
			Path:  path,
			Score: score,
		}

		rankings = append(rankings, ranking)
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Score > rankings[j].Score
	})

	for i := range rankings {
		results = append(
			results,
			// fmt.Sprintf("%s (%v)", rankings[i].Path, rankings[i].Score),
			rankings[i].Path,
		)
	}

	return results
}
