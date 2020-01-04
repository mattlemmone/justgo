package search

import (
  "fmt"
  "strings"
)

type Trie struct {
  root *trieNode
}

type trieNode struct {
  childrenByPrefix map[string]*trieNode
  value string
  end bool
}

func  (t *trieNode) children() []*trieNode {
  var childNodes []*trieNode

  for _, childNode := range t.childrenByPrefix {
    childNodes = append(childNodes, childNode)
  }

  return childNodes
}

func New() *Trie {
  trie := Trie{}
  trie.root = &trieNode{
    childrenByPrefix: map[string]*trieNode{},
  }

  return &trie
}

func (t *Trie) Insert(words []string){
  node := t.root
  
  for i := range words {
    word := words[i]

    if childNode, exists := node.childrenByPrefix[word]; exists {
      node = childNode
    } else {
      newChildNode := trieNode{
        value: word,
        childrenByPrefix: map[string]*trieNode{},
      }

      node.childrenByPrefix[word] = &newChildNode
      node = &newChildNode
    }
  }

  node.end = true
}

func (t *Trie) Search(targets ...string) []string {
  return searchFromNode(t.root, targets)
}

func searchFromNode(start *trieNode, targets []string) []string {
  var path strings.Builder

  node := start

  // get node at end of prefix match
  for i := range targets {
    word := targets[i]

    if childNode, exists := node.childrenByPrefix[word]; exists {
      path.WriteString(word)
      node = childNode
    } else {
      return []string{}
    }
  }   

  fmt.Printf("begin %s\n", path.String())
  return nodePathsToWords(node, &path)
}

// nodePathsToWords returns all words formed from children of "start" to leaves of the trie
func nodePathsToWords(start *trieNode, path *strings.Builder) []string{
  type pair struct {
    node *trieNode
    path string
  }

  var paths []string
  var nodeQueue []pair

  // queue start node
  nodeQueue = append(
    nodeQueue, 
    pair{
      node: start,
      path: path.String(),
  })

  // pop, update queue
  currentPair := &nodeQueue[0]
  nodeQueue = nodeQueue[1:]

  for currentPair != nil {
    if currentPair.node.end {
      fmt.Printf("add to paf %s\n", currentPair.path)
      paths = append(paths, currentPair.path)
    }

    fmt.Printf("the dam q %+v\n", nodeQueue)
    // queue all children
    for _, childNode := range currentPair.node.children() {
      fmt.Printf("child %+v\n", childNode.value)
      newPath := fmt.Sprintf("%s%s", currentPair.path, childNode.value)

      p := pair{
        node: childNode,
        path: newPath,
      }
      
      nodeQueue = append(nodeQueue, p)
    }

    if len(nodeQueue) == 0 {
      break
    }

    // pop, update queue
    currentPair = &nodeQueue[0]
    nodeQueue = nodeQueue[1:]
  }

  fmt.Printf("pathss %s\n", paths)
  return paths
}

