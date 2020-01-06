package search

func max(a int, b int) int {
	if a < b {
		return b 
	}

	return a
}

func min(a byte,b byte,c byte) byte {
	if a < b && a < c {
		return a 
	}

	if b < a && b < c {
		return b
  }

	return c
}
