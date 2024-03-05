// Implement Trie (Prefix Tree)
//
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
// New() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
//
// https://leetcode.com/problems/implement-trie-prefix-tree
// #trie
package trie

type Trie struct {
	root     *Trie
	data     rune
	terminal bool
	children [26]*Trie // 26 letters in alphabet
}

func New() *Trie {
	return &Trie{
		root: &Trie{},
	}
}

func (t *Trie) Insert(word string) {
	insert(t.root, word)
}

func insert(root *Trie, word string) {
	if len(word) == 0 {
		root.terminal = true
		return
	}
	i := word[0] - 'a'
	child := root.children[i]
	if child == nil {
		child = &Trie{data: rune(word[0])}
		root.children[i] = child
	}
	insert(child, word[1:])
}

func (t *Trie) Search(word string) bool {
	return search(t.root, word)
}

func search(root *Trie, word string) bool {
	if len(word) == 0 {
		return root.terminal
	}
	i := word[0] - 'a'
	child := root.children[i]
	return child != nil && search(child, word[1:])
}

func (t *Trie) StartsWith(prefix string) bool {
	return startsWith(t.root, prefix)
}

func startsWith(root *Trie, word string) bool {
	if len(word) == 0 {
		return true
	}
	i := word[0] - 'a'
	child := root.children[i]
	return child != nil && startsWith(child, word[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
