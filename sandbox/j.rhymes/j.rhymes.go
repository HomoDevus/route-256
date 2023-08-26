package main

import (
	"bufio"
	"fmt"
	"os"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	var dictLen, wordsAmount int
	dict := NewTrie()

	fmt.Fscan(in, &dictLen)

	for i := 0; i < dictLen; i++ {
		var line string

		fmt.Fscanf(in, "\n%s", &line)

		dict.Insert(line)
	}

	fmt.Fscan(in, &wordsAmount)

	for i := 0; i < wordsAmount; i++ {
		var word string

		fmt.Fscan(in, &word)

		fmt.Fprintln(out, dict.Suffics(word))
	}

	defer out.Flush()
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for i := len(word) - 1; i >= 0; i-- {
		char := rune(word[i])

		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}

		node = node.children[char]
	}

	node.isEnd = true
}

func (t *Trie) Suffics(word string) string {
	node := t.root
	lastFork := len(word)
	forkNode := t.root

	for i := len(word) - 1; i >= 0; i-- {
		char := rune(word[i])

		if _, exists := node.children[char]; !exists {
			return node.Word(word[i+1:], word)
		}

		node = node.children[char]

		if len(node.children) > 1 || (i == 0 && len(node.children) > 0) || (i != 0 && node.isEnd) {
			forkNode = node
			lastFork = i
		}
	}


	if lastFork == len(word) {
		return t.root.Word("", word)
	} else {
		return forkNode.Word(word[lastFork:], word)
	}

}

func (t *TrieNode) Word(suffics, targetWord string) string {
	node := t
	word := suffics


	for i := len(targetWord) - len(suffics) - 1; !node.isEnd || word == targetWord; i-- {
		for letter, tire := range node.children {
			if i < 0 || letter != rune(targetWord[i]) {
				word = string(letter) + word
				node = tire
				break
			}
		}
	}

	return word
}
