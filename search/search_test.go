package search

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestFuzzyFind(t *testing.T) {
    actualMatches := FuzzyFind(
      "target",
      "hey",
      "target",
      "ticket",
      "targets",
    )

    expectedMatches := []string {
      "target",
      "targets",
      "ticket",
      "hey",
    }

    assert.Equal(t, expectedMatches, actualMatches)
}

func TestFuzzyFindApplication(t *testing.T) {
    actualMatches := FuzzyFind(
      "Firefox",
      "Thunderbird",
      "Discord",
      "Alacritty",
    )

    expectedMatches := []string {
      "Discord",
      "Alacritty",
      "Thunderbird",
    }

    assert.Equal(t, expectedMatches, actualMatches)
}


// func TestFuzzyFindPath(t *testing.T) {
//     actualMatches := FuzzyFindFile(
//       "pictures",
//       "/home/matt/Documents/Pictures/",
//       "/home/matt/Pictures/Screenshot from 2020-01-01 17-56-09.png",
//       "/home/matt/Documents/Videos/",
//     )

//     expectedMatches := []string {
//       "/home/matt/Pictures/Screenshot from 2020-01-01 17-56-09.png",
//       "/home/matt/Documents/Pictures/",
//       "/home/matt/Documents/Videos/",
//     }

//     assert.Equal(t, expectedMatches, actualMatches)
// }

// func TestFuzzyFindFilePrefix(t *testing.T) {
//     actualMatches := FuzzyFindFile(
//       "home",
//       "/home/matt/Documents/Pictures/",
//       "/home/matt/Pictures/Screenshot from 2020-01-01 17-56-09.png",
//       "/home/matt/Documents/Videos/",
//     )

//     expectedMatches := []string {
//       "/home/matt/Documents/Videos/",
//       "/home/matt/Documents/Pictures/",
//       "/home/matt/Pictures/Screenshot from 2020-01-01 17-56-09.png",
//     }

//     assert.Equal(t, expectedMatches, actualMatches)
// }
