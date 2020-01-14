package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuzzyFindPathMatchingWord(t *testing.T) {
	actualMatches := fuzzyFindFile(
		"pictures",
		[]string{
			"/home/matt/Documents/Videos/pictures",
			"/home/matt/Documents/Pictures/",
			"/home/matt/Documents/Videos/pic",
			"/home/matt/Documents/Videos/",
			"/home/matt/Documents/Videos/pictures/pictures",
		},
	)

	expectedMatches := []string{
		"/home/matt/Documents/Pictures/",
		"/home/matt/Documents/Videos/pictures",
		"/home/matt/Documents/Videos/pictures/pictures",
	}

	assert.Equal(t, expectedMatches, actualMatches)
}

func TestFuzzyFindPathMatchingWordPrefix(t *testing.T) {
	actualMatches := fuzzyFindFile(
		"pic",
		[]string{
			"/home/matt/Documents/Videos/pictures",
			"/home/matt/Documents/Pictures/",
			"/home/matt/Documents/Videos/pic",
			"/home/matt/Documents/Videos/",
			"/home/matt/Documents/Videos/pictures/pictures",
		},
	)

	expectedMatches := []string{
		"/home/matt/Documents/Videos/pic",
		"/home/matt/Documents/Pictures/",
		"/home/matt/Documents/Videos/pictures",
		"/home/matt/Documents/Videos/pictures/pictures",
	}

	assert.Equal(t, expectedMatches, actualMatches)
}

func TestFuzzyFindPathMatchingWordSuffix(t *testing.T) {
	actualMatches := fuzzyFindFile(
		"tures",
		[]string{
			"/home/matt/Documents/Videos/pictures",
			"/home/matt/Documents/Pictures/",
			"/home/matt/Documents/Videos/pic",
			"/home/matt/Documents/Videos/",
			"/home/matt/Documents/Videos/pictures/pictures",
		},
	)

	expectedMatches := []string{
		"/home/matt/Documents/Pictures/",
		"/home/matt/Documents/Videos/pictures",
		"/home/matt/Documents/Videos/pictures/pictures",
	}

	assert.Equal(t, expectedMatches, actualMatches)
}
