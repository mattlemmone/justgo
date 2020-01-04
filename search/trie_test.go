package search

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestTrieInsertAndSearchFromFirstLevel(t *testing.T) {
    trie := New()

    trie.Insert([]string{
      "hello/",
      "everyone/",
      "as/",
      "well/",
      "as/",
      "the/",
      "world",
    })
    
    trie.Insert([]string{"goodbye/","cruel/", "world"})
    trie.Insert([]string{"hello/","everyone/"})
    trie.Insert([]string{"hello/","world"})
    trie.Insert([]string{"hello/"})
    trie.Insert([]string{"world"})
    
    actual := trie.Search("hello/")
    expected := []string{
      "hello/everyone/",
      "hello/everyone/as/well/as/the/world",
      "hello/world",
      "hello/",
    }

    assert.ElementsMatch(t, expected, actual)
}


func TestTrieInsertAndSearchFromFourthLevel(t *testing.T) {
    trie := New()

    trie.Insert([]string{"1", "2", "3"})
    trie.Insert([]string{"1", "2", "3", "4"})
    trie.Insert([]string{"1", "2", "3", "4", "5", "6", "7"})
    trie.Insert([]string{"1", "2", "3", "4", "5", "66", "77"})
    trie.Insert([]string{"1", "2", "3", "4", "5", "66", "77", "88"})
    trie.Insert([]string{"1", "2", "3", "4", "5", "66", "777", "888"})

    actual := trie.Search("1", "2", "3", "4")
    expected := []string{
      "1234",
      "1234567",
      "123456677",
      "12345667788",
      "1234566777888",
    }

    assert.ElementsMatch(t, expected, actual)
}
