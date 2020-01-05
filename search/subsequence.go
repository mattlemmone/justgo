package search

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

  return candidatePtr == len(candidate) - 1
}
