package algorithm

import (
	"fmt"
	"testing"
)

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor5() Trie {
	trie := Trie{}
	return trie
}

func (this *Trie) Insert(word string) {
	node := this
	runes := []rune(word)
	for _, rr := range runes {
		trie := node.children[rr-'a']
		if trie == nil {
			node.children[rr-'a'] = &Trie{}
		}
		node = node.children[rr-'a']
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	runes := []rune(prefix)
	node := this
	for _, r := range runes {
		trie := node.children[r-'a']
		if trie == nil {
			return nil
		}
		node = trie
	}

	return node
}

func Test208(t *testing.T) {

	//["Trie","insert","startsWith"]
	//[[],["hotdog"],["dog"]]
	trie := Constructor5()
	trie.Insert("hotdog")
	fmt.Println(trie.StartsWith("hot"))
}
