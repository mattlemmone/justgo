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
    // actual := subsequence("cat", "tact") 
    actual := subsequence(
      "n",
      "/home/matt/opt/ardour/gtk2_ardour/win32/msvc_resources.rc.in",
    ) 

    assert.False(t, actual)
}
