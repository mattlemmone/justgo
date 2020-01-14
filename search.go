package main

import (
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
	directoryBonusNumerator = 0.7
)

func fuzzyFindFile(target string, paths []string) []string {
	var rankings []FileRanking
	var results []string

	loweredTarget := strings.ToLower(target)

	for _, path := range paths {
		loweredPath := strings.ToLower(path)
		loweredPath = removeTrailingSlash(loweredPath)

		if !Subsequence(loweredTarget, loweredPath) {
			continue
		}

		rankings = append(rankings, FileRanking{
			Path:  path,
			Score: score(loweredTarget, loweredPath),
		})
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Score > rankings[j].Score
	})

	for i := range rankings {
		results = append(results, rankings[i].Path)
	}

	return results
}

func score(target string, path string) float64 {
	// todo: dont split, substr
	dirs := strings.Split(path, "/")[1:]

	// add bonus to prefer short paths
	score := directoryBonusNumerator / math.Log(float64(len(dirs)))
	lastFile := dirs[len(dirs)-1]

	return score + SubsequenceSimilarity(target, lastFile)
}

func removeTrailingSlash(s string) string {
	if s[len(s)-1] == '/' {
		return s[:len(s)-1]
	}

	return s
}
