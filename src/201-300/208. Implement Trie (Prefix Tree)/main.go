package main

import "fmt"

//type Trie struct {
//	val      [26] byte
//	end      [26] bool
//	children [26]*Trie
//}
//
///** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{}
//}
//
///** Inserts a word into the trie. */
//func (this *Trie) Insert(word string) {
//	if len(word) == 0 {
//		return
//	}
//	insert(this, word, 0, len(word))
//}
//func insert(this *Trie, word string, idx int, len1 int) {
//	temp := word[idx]
//	tempIdx := getIdx(temp)
//	this.val[tempIdx] = temp
//	if idx+1 < len1 {
//		next := getIdx(word[idx+1])
//		if this.children[next] == nil {
//
//			this.children[next] = new(Trie)
//		}
//		insert(this.children[next], word, idx+1, len1)
//	} else {
//		this.end[tempIdx] = true
//	}
//}
//
///** Returns if the word is in the trie. */
//func (this *Trie) Search(word string) bool {
//	if len(word) == 0 {
//		return true
//	}
//	return search(this, word, 0, len(word))
//}
//func search(this *Trie, word string, idx int, len1 int) bool {
//	temp := word[idx]
//	tempIdx := getIdx(temp)
//	if this.val[tempIdx] != temp {
//		return false
//	}
//	if idx+1 < len1 {
//		next := getIdx(word[idx+1])
//		if this.children[next] == nil {
//			return false
//		}
//		return search(this.children[next], word, idx+1, len1)
//	}
//	if this.end[tempIdx] == false {
//		return false
//	}
//	return true
//}
//
///** Returns if there is any word in the trie that starts with the given prefix. */
//func (this *Trie) StartsWith(prefix string) bool {
//	if len(prefix) == 0 {
//		return true
//	}
//	return startWith(this, prefix, 0, len(prefix))
//}
//func startWith(this *Trie, word string, idx int, len1 int) bool {
//	temp := word[idx]
//	tempIdx := getIdx(temp)
//	if this.val[tempIdx] != temp {
//		return false
//	}
//	if idx+1 < len1 {
//		next := getIdx(word[idx+1])
//		if this.children[next] == nil {
//			return false
//		}
//		return startWith(this.children[next], word, idx+1, len1)
//	}
//	return true
//}
//
//func getIdx(char byte) int {
//	return int(char - 'a')
//}

//the first 1 is header, do not store any val
type Trie1 struct {
	val  byte
	sons [26]*Trie1
	end  int
}

func Constructor1() Trie1 {
	return Trie1{}
}

func (this *Trie1) Insert1(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie1{val: word[i]}
		}

		node = node.sons[idx]
	}

	node.end++
}

func (this *Trie1) Search1(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

func (this *Trie1) StartsWith1(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := new(Trie1)
	trie.Insert1("aa")
	trie.Insert1("app")
	trie.Insert1("apple")
	trie.Insert1("beer")
	trie.Insert1("add")
	trie.Insert1("jam")
	trie.Insert1("rental")

	fmt.Println(trie.Search1("app"))
	fmt.Println(trie.Search1("apps"))
	fmt.Println(trie.Search1("ad"))
	fmt.Println(trie.Search1("applepie"))
	fmt.Println(trie.Search1("rest"))
	fmt.Println(trie.Search1("jan"))
	fmt.Println(trie.Search1("rent"))
	fmt.Println(trie.Search1("beer"))
	fmt.Println(trie.Search1("jam"))

	fmt.Println(trie.StartsWith1("apps"))
	fmt.Println(trie.StartsWith1("app"))
	fmt.Println(trie.StartsWith1("ad"))
	fmt.Println(trie.StartsWith1("applepie"))
	fmt.Println(trie.StartsWith1("rest"))
	fmt.Println(trie.StartsWith1("jan"))
	fmt.Println(trie.StartsWith1("rent"))
	fmt.Println(trie.StartsWith1("beer"))
	fmt.Println(trie.StartsWith1("jam"))

}
