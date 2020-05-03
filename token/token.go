package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NEQ      = "!="

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywores
	FUNCTION = "FUNCTION"
	RETURN   = "return"
	LET      = "LET"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
)

func LookupIdent(ident string) TokenType {
	var keywords = map[string]TokenType{
		"fn":     FUNCTION,
		"return": RETURN,
		"let":    LET,
		"if":     IF,
		"else":   ELSE,
		"true":   TRUE,
		"false":  FALSE,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
