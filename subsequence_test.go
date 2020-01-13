package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsequenceSame(t *testing.T) {
	actual := Subsequence("cat", "cat")

	assert.True(t, actual)
}

func TestSubsequenceSimple(t *testing.T) {
	actual := Subsequence("cat", "ratcatbrat")

	assert.True(t, actual)
}

func TestSubsequenceSimpleSplit(t *testing.T) {
	actual := Subsequence("cat", "c_a_t")

	assert.True(t, actual)
}

func TestInvalidSubsequenceSimple(t *testing.T) {
	actual := Subsequence("cat", "rat")

	assert.False(t, actual)
}

func TestInvalidSubsequenceRearranged(t *testing.T) {
	actual := Subsequence("cat", "tact")

	assert.False(t, actual)
}

func TestSubsequenceSimilaritySame(t *testing.T) {
	actual := SubsequenceSimilarity("cat", "cat")
	expected := 1.0

	assert.Equal(t, expected, actual)
}

func TestSubsequenceSimilarityGapWeight(t *testing.T) {
	case1 := SubsequenceSimilarity("disc", "discord")
	case2 := SubsequenceSimilarity("diso", "discord")

	assert.Greater(t, case1, case2)
}

func TestSubsequenceSimilarityGapWeight2(t *testing.T) {
	case1 := SubsequenceSimilarity("diso", "discord")
	case2 := SubsequenceSimilarity("dsrd", "discord")

	assert.Greater(t, case1, case2)
}

func TestSubsequenceSimilarityGreediness(t *testing.T) {
	case1 := SubsequenceSimilarity("disc", "discord")
	case2 := SubsequenceSimilarity("isco", "discord")

	assert.Greater(t, case1, case2)
}

func TestSubsequenceSimilarityShortText(t *testing.T) {
	case1 := SubsequenceSimilarity("disco", "d")

	assert.Less(t, case1, 1.0)
}

func TestSubsequenceSimilarityMostlySimilar(t *testing.T) {
	case1 := SubsequenceSimilarity("py3", "python3.7")
	case2 := SubsequenceSimilarity("py3", "python2.6")

	assert.Greater(t, case1, case2)
}
