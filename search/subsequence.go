package search

const (
  // significance of matching subsequence at start
  initialGapWeight       = float64(0.20)

  // significance of having subsequence grouped closely together
  contiguityWeight       = float64(0.45)

  // significance of matching entire string
  sequenceCoverageWeight = float64(0.35)
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

      gapLen += (i - lastCandidate - 1)
      lastCandidate = i 
      candidatePtr++
    }
  }

  maxLen := float64(max(len(candidate), len(text)))
  coverageScore := float64(candidatePtr) / maxLen
  contiguityScore := 1 - float64(gapLen) / maxLen
  initialGapScore := 0.0

  if initialGap != -1 {
    initialGapScore = 1 - float64(initialGap) / maxLen
  }

  return coverageScore * sequenceCoverageWeight +
    contiguityScore * contiguityWeight +
    initialGapScore * initialGapWeight 
}
