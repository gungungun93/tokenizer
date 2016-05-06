package tokenizer

import (
	"unicode"
	"strings"
)

//-----------------------------------------------------------------------------
// LongLexto Constants
//-----------------------------------------------------------------------------
// 1. A default file name to dictionary for constructing a Trie
// 2. Types of tokenized text: Latin, Cyrilic, Number, Whitespace, Thai, etc...
//-----------------------------------------------------------------------------
const (
	SPACE = "SPACE"
	NUMBER = "NUMBER"
	WESTERN = "WESTERN"
	THAI = "THAI"
	PUNC = "SPECIAL"
	TAG = "TAG"
)
//-----------------------------------------------------------------------------
// LongLexto Structure
//-----------------------------------------------------------------------------
// A tokenizer with an ability to tokenize Thai language
//-----------------------------------------------------------------------------
type LongLexto struct {
	pTree LongParseTree
	tokens []Token
	current int
	text string
}
//-----------------------------------------------------------------------------
// LongLexto Methods
//-----------------------------------------------------------------------------
// 1. GetText:
// Returns the original text untokenized
func (lexto *LongLexto) GetText() string {
	return lexto.text
}
//-----------------------------------------------------------------------------
// 2. First:
// Returns the first token
func (lexto *LongLexto) First() Token {
	return lexto.tokens[0]
}
//-----------------------------------------------------------------------------
// 3. Next:
// Returns the next token, or an empty string, if end of token list
func (lexto *LongLexto) Next() Token {
	if lexto.HasNext() {
		result := lexto.tokens[lexto.current]
		lexto.current += 1
		return result
	}
	return Token{}
}
//-----------------------------------------------------------------------------
// 4. Previous:
// Returns the next token, or an empty string, if at the beginning of token list
func (lexto *LongLexto) Previous() Token {
	if lexto.HasPrevious() {
		result := lexto.tokens[lexto.current]
		lexto.current -= 1
		return result
	}
	return Token{}
}
//-----------------------------------------------------------------------------
// 5. HasNext:
// Returns true if the current token is not the last element in the token list
func (lexto LongLexto) HasNext() bool {
	return lexto.current < len(lexto.tokens)
}
//-----------------------------------------------------------------------------
// 6. HasPrevious:
// Returns true if the current token is not the first element in the token list
func (lexto LongLexto) HasPrevious() bool {
	return lexto.current > -1
}
//-----------------------------------------------------------------------------
// 7. SetText:
// Assigns new text and tokenizes the text completely
// Results in the list of token to be ready for use
func (lexto *LongLexto) SetText(newStr string) {
	// 1. Empty strings not accepted!
	if len(newStr) == 0 {
		return
	}

	// 2. Reset token list and reassign new text
	lexto.tokens = []Token{}
	lexto.text = newStr

	// 3. Read each character in the string
	var i int = 0
	var j int
	var chars []rune = []rune(newStr)

	for i < len(chars) {
		// 4. Check if the character is a whitespace
		if unicode.IsSpace(chars[i]) {
			// 4.1 Keep reading until the character is no longer a white space
			j = i + 1
			for j < len(chars) && unicode.IsSpace(chars[j]) {
				j += 1
			}
			// 4.2 Append the character as whitespace type token
			lexto.tokens = append(lexto.tokens, Token{text : string(chars[i : j]), textType : SPACE})
			i = j
		// 5. Check if the character is Latin, Cyrillic, or Greek
		} else if IsWestern(chars[i]) {
			// 5.1 Keep reading until the character is no longer from Western language
			j = i + 1
			for j < len(chars) && IsWestern(chars[j]) {
				j += 1
			}
			// 5.2 Convert all letters to lower case while adding tokens to the list
			lexto.tokens = append(lexto.tokens, Token{text : strings.ToLower(string(chars[i : j])), textType : WESTERN})
			i = j
		// 6. Check if the character is a digit
		} else if unicode.IsNumber(chars[i]) {
			// 6.1 Keep reading until the character is no longer digits
			j = i + 1
			for j < len(chars) && unicode.IsNumber(chars[j]) {
				j += 1
			}
			// 6.2 Append the character as digits type token
			lexto.tokens = append(lexto.tokens, Token{text : string(chars[i : j]), textType : NUMBER})
			i = j
		// 7. Check if the character is Thai
		} else if IsThai(chars[i]) {
			// 7.1 Obtain the position of next word
			j = lexto.pTree.parseWordInstance(i, chars, lexto.tokens)
			// 7.2 Append the word to the list as Thai word
			lexto.tokens = append(lexto.tokens, Token{text : string(chars[i : j]), textType : THAI})
			i = j
		// 8. Check if the character is a HTML Tag
		} else if chars[i] == '<' {
			// 8.1 Keep reading until the character is '>'
			j = i + 1
			for j < len(chars) && chars[j] != '>' {
				j += 1
			}
			// 8.2 Append the character as HTML tag type token
			lexto.tokens = append(lexto.tokens, Token{text : string(chars[i : j + 1]), textType : TAG})
			i = j + 1
		// 9. Check if the character is a symbol or punctuation
		} else {
			// 9.1 Just simply add the symbol and move on to the next character
			lexto.tokens = append(lexto.tokens, Token{text : string(chars[i]), textType : PUNC})
			i += 1
		}
	}

	lexto.current = 0
}
//-----------------------------------------------------------------------------
// Tokenizer Functions
//-----------------------------------------------------------------------------
// 1. IsWestern:
// Returns true, if the utf8 of the character is either Latin, Greek, or Cyrilic
func IsWestern(char rune) bool {
	return (char > '\u0040' && char < '\u005B') ||
		   (char > '\u0060' && char < '\u007B') ||
		   (char > '\u00BF' && char < '\u00D7') ||
		   (char > '\u00D7' && char < '\u02B0') ||
		   (char > '\u036F' && char < '\u0590')
}
//-----------------------------------------------------------------------------
// 2. IsThai:
// Returns true, if the utf8 of the character is a Thai script
func IsThai(char rune) bool {
	return (char > '\u0E00' && char < '\u0E3B') ||
		   (char > '\u0E3F' && char < '\u0E5C')
}
//-----------------------------------------------------------------------------
// 3. Initialize:
// Returns a Thai tokenizer, given a dictionary file
func Initialize(file string) *LongLexto {
	return &LongLexto{current : 0, text : "", pTree : CreateTree(file), tokens : []Token{}}
}
//-----------------------------------------------------------------------------
// 4. Tokenize: For Special Ocassions
// Returns a list of tokens obtained from tokenizing texts, given a text string
// and a tokenizer
func Tokenize(tokenizer *LongLexto, text string) []Token {
	tokenizer.SetText(text)
	result := []Token{}

	for tokenizer.HasNext() {
		result = append(result, tokenizer.Next())
	}

	return result
}