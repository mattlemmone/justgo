package search


func levenshteinDistance(s1 string, s2 string) float64 {
  if s1 == s2 {
    return 1.0
  }

  n := len(s1) 
  m := len(s2)
    
  matrix := make([][]byte, n + 1)
  rows := make([]byte, (n + 1) * (m + 1))

  for i := range matrix {
    matrix[i] = rows[i * (m + 1) : (i+1) * (m+1)]
  }

  for i := range matrix {
    matrix[i][0] = byte(i)
  }

  for i := range matrix[0] {
    matrix[0][i] = byte(i)
  }

  for i := 1; i <= n ; i++ {
    for j:= 1; j <= m; j++ {
      if s1[i - 1] == s2[j - 1] {
        matrix[i][j] = matrix[i - 1][j - 1]
        continue
      }
			
			insertion   := matrix[i][j - 1];
			deletion    := matrix[i - 1][j];
			replacement := matrix[i - 1][j - 1];

			matrix[i][j] =  1 + min(insertion, deletion, replacement);	
    }
  }

  edits := matrix[n][m]

  return float64(edits) / float64(max(n,m)) 
}

