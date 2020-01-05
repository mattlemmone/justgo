package search

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSubsequenceSame(t *testing.T) {
    actual := subsequence("cat", "cat") 

    assert.True(t, actual)
}


func TestSubsequenceSimple(t *testing.T) {
    actual := subsequence("cat", "ratcatbrat") 

    assert.True(t, actual)
}

func TestSubsequenceSimpleSplit(t *testing.T) {
    actual := subsequence("cat", "c_a_t") 

    assert.True(t, actual)
}

func TestInvalidSubsequenceSimple(t *testing.T) {
    actual := subsequence("cat", "rat") 

    assert.False(t, actual)
}

func TestInvalidSubsequenceRearranged(t *testing.T) {
    actual := subsequence("cat", "tact") 

    assert.False(t, actual)
}

func TestSubsequenceSimilaritySame(t *testing.T) {
    actual := subsequenceSimilarity("cat", "cat") 
    expected := 1.0

    assert.Equal(t,  expected, actual)
}

func TestSubsequenceSimilarityGapWeight(t *testing.T) {
    case1 := subsequenceSimilarity("disc", "discord") 
    case2 := subsequenceSimilarity("diso", "discord") 

    assert.Greater(t, case1, case2)
}

func TestSubsequenceSimilarityGapWeight2(t *testing.T) {
    case1 := subsequenceSimilarity("diso", "discord") 
    case2 := subsequenceSimilarity("dsrd", "discord") 

    assert.Greater(t, case1, case2)
}

func TestSubsequenceSimilarityGreediness(t *testing.T) {
    case1 := subsequenceSimilarity("disc", "discord") 
    case2 := subsequenceSimilarity("isco", "discord") 

    assert.Greater(t, case1, case2)
}


func TestSubsequenceSimilarityShortText(t *testing.T) {
    case1 := subsequenceSimilarity("disco", "d") 

    assert.Less(t, case1, 1.0)
}
