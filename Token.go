package tokenizer

//-----------------------------------------------------------------------------
// Token Structure
//-----------------------------------------------------------------------------
// A token structure which stores a string token and its type
//-----------------------------------------------------------------------------
type Token struct {
    text string
    textType string
}
//-----------------------------------------------------------------------------
// Token Methods
//-----------------------------------------------------------------------------
// 1. GetText:
// Returns a text content of the token
func (token Token) GetText() string {
    return token.text
}
//-----------------------------------------------------------------------------
// 2. GetTextType:
// Returns a text type of the token
func (token Token) GetTextType() string {
    return token.textType
}
//-----------------------------------------------------------------------------
// 3. IsSpace:
// Returns true, if the token is a type of whitespace
func (token Token) IsSpace() bool {
    return token.textType == SPACE
}
//-----------------------------------------------------------------------------
// 4. IsNumber:
// Returns true, if the token is a type of numbers
func (token Token) IsNumber() bool {
    return token.textType == NUMBER
}
//-----------------------------------------------------------------------------
// 5. IsNumber:
// Returns true, if the token is a type of punctuations
func (token Token) IsSymbol() bool {
    return token.textType == PUNC
}
//-----------------------------------------------------------------------------
// 5. IsNumber:
// Returns true, if the token is a type of punctuations
func (token Token) IsHTML() bool {
    return token.textType == TAG
}