package tokenizer

//-----------------------------------------------------------------------------
// LongParseTree Structure
//-----------------------------------------------------------------------------
// A tree structure for storing chains of characters
// which forms a word
//-----------------------------------------------------------------------------
type LongParseTree struct {
	dictionary Trie
	frontDepChar map[rune]bool
	rearDepChar map[rune]bool
	tonalChar map[rune]bool
	endingChar map[rune]bool
}
//-----------------------------------------------------------------------------
// LongParseTree Methods
//-----------------------------------------------------------------------------
// 1. nextWordValid: Package Internal Use
// Finds the longest chain of characters from received string, which does
// exist in a Trie dictionary and is considered a valid word.
// Returns a position which marks the end of longest valid word.
func (pTree LongParseTree) nextWordValid(beginPos int, text []rune) bool {
	var pos int = beginPos + 1
	var status int

	if beginPos == len(text) {
		return true
	} else if text[beginPos] <= '~' {
		return true
	} else {
		for pos <= len(text) {
			status = pTree.dictionary.Contains(string(text[beginPos :  pos]))

			if status == 1 {
				return true
			} else if status == 0 {
				pos += 1
			} else {
				break
			}
		}
	}

	return false
}
//-----------------------------------------------------------------------------
// 2. parseWordInstance: Package Internal Use
// Finds the longest chain of characters from received string, which does
// exist in a Trie dictionary.
// Returns a position which marks the longest word.
func (pTree LongParseTree) parseWordInstance(beginPos int, text []rune, list []Token) int {
	var longestPos, longestValidPos, returnPos int = -1, -1, -1
	var pos, status int = beginPos + 1, 1
	var prevChar rune = '0'

	for pos <= len(text) && status != -1 {
		status = pTree.dictionary.Contains(string(text[beginPos :  pos]))

		if status == 1 {
			longestPos = pos

			if pTree.nextWordValid(pos, text) {
				longestValidPos = pos
			}
		}

		pos += 1
	}

	if(beginPos >= 1) {
		prevChar = text[beginPos - 1]
	}

	if longestPos == -1 {
		returnPos = beginPos + 1

		_, found := pTree.frontDepChar[text[beginPos]]
		_, found2 := pTree.tonalChar[text[beginPos]]
		_, found3 := pTree.rearDepChar[prevChar]

		if len(list) > 0 && (found || found2 || found3) {
			list[len(list) - 1].text = string(text[beginPos : returnPos])
		} else {
			list = append(list, Token{text : string(text[beginPos : returnPos]), textType : THAI})
		}
		return returnPos
	} else {
		if longestValidPos == -1 {
			if _, found := pTree.rearDepChar[prevChar]; found {
				list[len(list) - 1].text = string(text[beginPos : longestPos])
			} else {
				list = append(list, Token{text : string(text[beginPos : longestPos]), textType : THAI})
			}
			return longestPos
		} else {
			if _, found := pTree.rearDepChar[prevChar]; found {
				list[len(list) - 1].text = string(text[beginPos : longestValidPos])
			} else {
				list = append(list, Token{text : string(text[beginPos : longestValidPos]), textType : THAI})
			}
			return longestValidPos
		}
	}
}
//-----------------------------------------------------------------------------
// Constructor function
//-----------------------------------------------------------------------------
// CreateTree:
// Creates a parse tree for tokenizing Thai words using LM algorithm.
func CreateTree(file string) LongParseTree {
	pTree := LongParseTree{dictionary : CreateTrie(file)}

	pTree.frontDepChar = make(map[rune]bool)
	pTree.frontDepChar['Ð'] = true
	pTree.frontDepChar['Ñ'] = true
	pTree.frontDepChar['Ò'] = true
	pTree.frontDepChar['Ó'] = true
	pTree.frontDepChar['Ô'] = true
	pTree.frontDepChar['Õ'] = true
	pTree.frontDepChar['Ö'] = true
	pTree.frontDepChar['×'] = true
	pTree.frontDepChar['Ø'] = true
	pTree.frontDepChar['Ù'] = true
	pTree.frontDepChar['å'] = true
	pTree.frontDepChar['ç'] = true
	pTree.frontDepChar['ì'] = true
	pTree.frontDepChar['í'] = true

	pTree.rearDepChar = make(map[rune]bool)
	pTree.rearDepChar['Ñ'] = true
	pTree.rearDepChar['×'] = true
	pTree.rearDepChar['à'] = true
	pTree.rearDepChar['á'] = true
	pTree.rearDepChar['â'] = true
	pTree.rearDepChar['ã'] = true
	pTree.rearDepChar['ä'] = true
	pTree.rearDepChar['í'] = true

	pTree.tonalChar = make(map[rune]bool)
	pTree.tonalChar['è'] = true
	pTree.tonalChar['é'] = true
	pTree.tonalChar['ê'] = true
	pTree.tonalChar['ë'] = true

	pTree.endingChar = make(map[rune]bool)
	pTree.endingChar['æ'] = true
	pTree.endingChar['Ï'] = true

	return pTree
}