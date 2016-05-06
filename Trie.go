package tokenizer

import (
	"bufio"
	"os"
)

const (
	DICT_FILE = "lexitron.txt"
)

//-----------------------------------------------------------------------------
// Trie Structure
//-----------------------------------------------------------------------------
// A tree-like structure for storing chains of characters
// which forms a word
//-----------------------------------------------------------------------------
type Trie struct {
	children map[rune]*Trie
	character rune
	isWord bool
}
//-----------------------------------------------------------------------------
// Trie Methods
//-----------------------------------------------------------------------------
// 1. AddWord:
// Adds a word to a trie as a chain of characters
func (trie Trie) AddWord(word string) {
	iter := &trie

	// 1. Read each character in a word
	for _, character := range word {
		// 2. Check if the character matches any children
		// 2.1 Skip all white spaces
		if value, found := iter.children[character]; found {
			iter = value
		// 2.3 If not match - Create a new node instead
		} else {
			// 2.3 Create a new node
			newNode := Trie{make(map[rune]*Trie), character, false}
			// 2.3 Add this new node as a new child
			iter.children[character] = &newNode
			// 2.4 Go to the node
			iter = iter.children[character]
		}
	}

	// 3. Set the leaf node with the last character in the word as a valid word
	iter.isWord = true
}
//-----------------------------------------------------------------------------
// 2. Contains:
// Tells whether a given word in a Trie is a valid word
// Returns -1 if the word is not in a dictionary
// Returns 0 if the word is in a dictionary, but not valid
// Returns 1 if the word is in a dictionary, and is a word
func (trie Trie) Contains(word string) int {
	iter := &trie

	// 1. Read each character in a word
	for _, character := range word {
		// 2. Check if the character matches any children
		if value, found := iter.children[character]; found {
			// 2.1 If match - Go to that node
			iter = value
		// 3. If there is no match the there exist no such node
		} else {
			return -1
		}
	}

	if iter.isWord {
		return 1
	} else {
		return 0
	}
}
//-----------------------------------------------------------------------------
// Constructor function
//-----------------------------------------------------------------------------
// 3. CreateTrie:
// Creates a dictionary trie, provided a dictionary text file
// Returns an empty trie, if the dictionary file does not exist
func CreateTrie(filename string) Trie {
	// 1. Create a dictionary structure
	dict := Trie{children : make(map[rune]*Trie)}
	// 2. Attempt to open the file
	file, err := os.Open(filename)

	// 3. If file not found, use a default file
	if err != nil {
		file, err = os.Open(DICT_FILE)
		// 3.1 If problem persists, return an empty trie
		if err != nil {
			return Trie{}
		}
	}
	defer file.Close()

	// 4. Read each word in a dictionary file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 5. Add the word to the trie
		dict.AddWord(scanner.Text())
	}

	return dict
}
