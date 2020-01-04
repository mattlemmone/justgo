package search

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestLevenshteinSimilar(t *testing.T) {
    actual := levenshteinDistance("cat", "hat") 
    expected := 1.0

    assert.Equal(t, expected, actual)
}


func TestNonTrivialCase(t *testing.T) {
    actual := levenshteinDistance("magical", "man, radical!") 
    expected := 0.14285714285714285

    assert.Equal(t, expected, actual)
}
