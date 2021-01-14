package antlr4

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/newm4n/symon-agen/antlr4/symonlang"
	"testing"
)

func TestLexer_test(t *testing.T) {
	data := "FooBar(234, 32423, Sub(\"3453\" + 3242))"
	is := antlr.NewInputStream(string(data))

	// Create the Lexer
	lexer := Symonlang.NewsymonlangLexer(is)
	//lexer := parser.NewLdifParserLexer(is)

	// Read all tokens
	for {
		nt := lexer.NextToken()
		fmt.Println(nt.GetText())
		if nt.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}
