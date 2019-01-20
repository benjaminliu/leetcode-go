package main

import "fmt"

type TrieNode struct {
	end      bool
	children [26]*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: new(TrieNode)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	addWord(this.root, word, 0, len(word))
}
func addWord(this *TrieNode, word string, idx int, len1 int) {
	if idx >= len1 {
		return
	}
	c := word[idx]
	i := c2i(c)
	if this.children[i] == nil {
		this.children[i] = &TrieNode{}
	}
	addWord(this.children[i], word, idx+1, len1)

	if idx == len1-1 {
		this.children[i].end = true
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return true
	}

	return search(this.root, word, 0, len(word))
}
func search(this *TrieNode, word string, idx int, len1 int) bool {
	if idx >= len1 {
		return true
	}

	c := word[idx]

	if c != '.' {
		i := c2i(c)
		if this.children[i] == nil {
			return false
		}
		if idx == len1-1 {
			if this.children[i].end == true {
				return true
			}
			return false
		}

		return search(this.children[i], word, idx+1, len1)
	} else {
		for i := 0; i < 26; i++ {
			if this.children[i] != nil {
				if idx == len1-1 {
					if this.children[i].end == true {
						return true
					}
				} else {
					if (search(this.children[i], word, idx+1, len1)) {
						return true
					}
				}
			}
		}
		return false
	}
}

func c2i(c byte) int {
	return int(c - 'a')
}

func main() {
	obj := Constructor();
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")
	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search(".at"))
	obj.AddWord("bat")
	fmt.Println(obj.Search(".at"))
	fmt.Println(obj.Search("an."))
	fmt.Println(obj.Search("a.d."))
	fmt.Println(obj.Search("b."))
	fmt.Println(obj.Search("a.d"))
	fmt.Println(obj.Search("."))
}
