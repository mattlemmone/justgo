package search

import "fmt"

const (
  initialGapWeight       = float64(0.45)
  contiguityWeight       = float64(0.35)
  sequenceCoverageWeight = float64(0.2)
)

func subsequence(candidate string, text string) bool {
  candidatePtr := 0

  for _, char := range text {
    if candidatePtr >= len(candidate) {
      break
    }

    if byte(char) == candidate[candidatePtr] {
      candidatePtr++
    }
  }

  return candidatePtr == len(candidate) 
}

func subsequenceSimilarity(candidate string, text string) float64 {
  candidatePtr := 0
  initialGap := -1
  lastCandidate := -1
  gapLen := 0

  for i, char := range text {
    if candidatePtr >= len(candidate) {
      break
    }

    if byte(char) == candidate[candidatePtr] {
      if initialGap == - 1 {
        initialGap = i
      }

      fmt.Printf("i - last: %v - %v\n", i, lastCandidate)
      gapLen += (i - lastCandidate - 1)
      lastCandidate = i 
      candidatePtr++
    }
  }

  coverageScore := float64(candidatePtr) / float64(len(text)) 
  contiguityScore := 1 - float64(gapLen) / float64(len(text))
  initialGapScore := 0.0

  if initialGap != -1 {
    initialGapScore = 1 - float64(initialGap) / float64(len(text))
  }

  fmt.Printf("%s in %s: cov %v, cont %v, gap %v\n", candidate, text, coverageScore, contiguityScore, initialGapScore)

  return coverageScore * sequenceCoverageWeight +
    contiguityScore * contiguityWeight +
    initialGapScore * initialGapWeight 
}
