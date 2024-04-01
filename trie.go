// You can edit this code!
// Click here and start typing.
package main

import (
	"log"
)

//current supported characters in node are characters with ASCII(32 - 126)
type Node struct {
	Char string
	Children    [95]*Node
	isEndOfWord bool
	isRoot      bool
	TableCols   interface{}
}

//32 offset to ignore the characters with ASCII(0-31)
const offset int = 32
const minIndex int = 0
const maxIndex int = 95
// Trie Root node
type Trie struct {
	TableName string
	RootNode  *Node
}

// NewTrie Creates a new trie
func NewTrie(mtName string) *Trie {
	//create root node
	root := NewNode("\000")
	root.isRoot = true
	return &Trie{TableName: mtName, RootNode: root}
}
func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := minIndex; i < maxIndex; i++ {
		node.Children[i] = nil
	}
	return node
}

func (t *Trie) Insert(word string, mapToAppend interface{}) {
	current := t.RootNode
	for i := minIndex; i < len(word); i++ {
		index := int(word[i]) - offset
		if index < minIndex || index >= maxIndex {
			break
		}
		if current == nil || current.Children[index] == nil {
			current.Children[index] = NewNode(string(word[i]))
		}
		current = current.Children[index]
	}
	current.TableCols = mapToAppend
	current.isEndOfWord = true
}

// Search will return false if a word we are searching for is not in the trie
func (t *Trie) Search(word string) (interface{}, bool) {
	current := t.RootNode
	for i := minIndex; i < len(word); i++ {
		index := int(word[i]) - offset
		if current == nil || current.Children[index] == nil {
			if current.isRoot {
				return nil, false
			} else {
				return current.TableCols, current.isEndOfWord
			}
		}
		current = current.Children[index]
	}
	return current.TableCols, current.isEndOfWord
}

func main() {
  //Sample implementation
	spCharac := " Test19\\|`,_+@?<>!#$%^&*()-=/.:;\"'{}[]~"
	t := NewTrie("TEST_TABLE")
	t.Insert("ME", []string{"m", "e"})
	t.Insert("MT", []string{"m", "t"})
	t.Insert(spCharac, []string{"s", "p"})
	t.SearchNode("MR")
	t.SearchNode("MER")
	t.SearchNode("M")
	t.SearchNode("MT")
	t.SearchNode("MTRA")
	t.SearchNode(spCharac + "abcd")
}

func (t *Trie) SearchNode(prefixToSearch string) {
	if v, ok := t.Search(prefixToSearch); ok {
		log.Println(prefixToSearch, v)
	}

}
