package antlr4

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/newm4n/symon-agen/antlr4/symonlang"
	"reflect"
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

func TestParse_test (t *testing.T) {
	data := "func(1234 + 5678)"
	sdata := string(data)

	is := antlr.NewInputStream(sdata)
	lexer := Symonlang.NewsymonlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := NewSymonAgentListener()

	psr := Symonlang.NewsymonlangParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Expression())
}


func TestExpression_test (t *testing.T) {
	data := "1234 + 4321 - 1111"
	sdata := string(data)

	is := antlr.NewInputStream(sdata)
	lexer := Symonlang.NewsymonlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := NewSymonAgentListener()

	psr := Symonlang.NewsymonlangParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Expression())

	rootExpr := listener.Root
	interf, err := rootExpr.Evaluate()
	if err != nil {
		t.Errorf(err.Error())
	}
	if reflect.ValueOf(interf).Kind() != reflect.Int64 {
		t.Errorf("not an int64")
	}
	if interf.(int64) != 4444 {
		t.Errorf("incorrect result")
	}
}


func TestFunctionExpression_test (t *testing.T) {
	data := "ToUpper(\"ThisIsAString\" + 1234)"
	sdata := string(data)

	is := antlr.NewInputStream(sdata)
	lexer := Symonlang.NewsymonlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := NewSymonAgentListener()

	psr := Symonlang.NewsymonlangParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Expression())

	rootExpr := listener.Root
	interf, err := rootExpr.Evaluate()
	if err != nil {
		t.Errorf(err.Error())
	}
	if reflect.ValueOf(interf).Kind() != reflect.String {
		t.Errorf("not an int64")
	}
	if interf.(string) != "THISISASTRING1234" {
		t.Errorf("incorrect result")
	}
}