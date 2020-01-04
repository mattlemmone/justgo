package search

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestTrivialCase(t *testing.T) {
    actualDistance := levenshteinDistance("cat", "hat") 
    expectedDistance := byte(1)

    assert.Equal(t, expectedDistance, actualDistance)
}


func TestNonTrivialCase(t *testing.T) {
    actualDistance := levenshteinDistance("magical", "man, radical!") 
    expectedDistance := byte(7)

    assert.Equal(t, expectedDistance, actualDistance)
}
