package antlr4

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/newm4n/symon-agen/antlr4/ast"
	Symonlang "github.com/newm4n/symon-agen/antlr4/symonlang"
	"log"
	"reflect"
	"strconv"
)

func NewSymonAgentListener() *SymonAgentListener {
	return &SymonAgentListener{
		Stack:                 newStack(),
		Root:                  nil,
	}
}

type SymonAgentListener struct {
	Symonlang.BasesymonlangListener
	Stack         *stack
	Root *ast.Expression
}

// EnterEveryRule is called when any rule is entered.
func (s *SymonAgentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *SymonAgentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *SymonAgentListener) EnterExpression(ctx *Symonlang.ExpressionContext) {
	expression := &ast.Expression{}
	s.Stack.Push(expression)
}

// ExitExpression is called when production expression is exited.
func (s *SymonAgentListener) ExitExpression(ctx *Symonlang.ExpressionContext) {
	expression := s.Stack.Pop().(*ast.Expression)
	if s.Stack.length == 0 {
		s.Root = expression
	} else {
		parent := s.Stack.Peek()
		if reflect.ValueOf(parent).Type().String() == "*ast.Expression" {
			parent.(*ast.Expression).AcceptChildExpression(expression)
		} else if reflect.ValueOf(parent).Type().String() == "*ast.ArgumentList" {
			parent.(*ast.ArgumentList).AcceptChildExpression(expression)
		} else {
			log.Fatalf("Parent is not an expression nor argumentList. Its a %s", reflect.ValueOf(parent).Type().String())
		}
	}
}

// EnterFunction is called when production function is entered.
func (s *SymonAgentListener) EnterFunction(ctx *Symonlang.FunctionContext) {
	function := &ast.Function{
		Name: ctx.FUNCTION_NAME().GetText(),
	}
	s.Stack.Push(function)
}

// ExitFunction is called when production function is exited.
func (s *SymonAgentListener) ExitFunction(ctx *Symonlang.FunctionContext) {
	function := s.Stack.Pop().(*ast.Function)
	parent := s.Stack.Peek().(*ast.Expression)
	parent.AcceptFunction(function)
}

// EnterConstant is called when production constant is entered.
func (s *SymonAgentListener) EnterConstant(ctx *Symonlang.ConstantContext) {
	constant := &ast.Constant{}
	s.Stack.Push(constant)
}

// ExitConstant is called when production constant is exited.
func (s *SymonAgentListener) ExitConstant(ctx *Symonlang.ConstantContext) {
	constant := s.Stack.Pop().(*ast.Constant)
	parent := s.Stack.Peek().(*ast.Expression)
	parent.AcceptConstant(constant)
}

// EnterOperator is called when production operator is entered.
func (s *SymonAgentListener) EnterOperator(ctx *Symonlang.OperatorContext) {
	operator := &ast.Operator{}
	operator.Operator = ctx.GetText()
	s.Stack.Push(operator)
}

// ExitOperator is called when production operator is exited.
func (s *SymonAgentListener) ExitOperator(ctx *Symonlang.OperatorContext) {
	operator := s.Stack.Pop().(*ast.Operator)
	parent := s.Stack.Peek().(*ast.Expression)
	parent.AcceptOperator(operator)
}

// EnterArgumentList is called when production argumentList is entered.
func (s *SymonAgentListener) EnterArgumentList(ctx *Symonlang.ArgumentListContext) {
	argList := &ast.ArgumentList{}
	s.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (s *SymonAgentListener) ExitArgumentList(ctx *Symonlang.ArgumentListContext) {
	argList := s.Stack.Pop().(*ast.ArgumentList)
	parent := s.Stack.Peek().(*ast.Function)
	parent.Args = argList
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *SymonAgentListener) EnterStringLiteral(ctx *Symonlang.StringLiteralContext) {
	stringValu := ctx.GetText()[1:len(ctx.GetText())-1]
	sLiteral := &ast.StringLiteral{
		Value: stringValu,
	}
	s.Stack.Push(sLiteral)
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *SymonAgentListener) ExitStringLiteral(ctx *Symonlang.StringLiteralContext) {
	strLit := s.Stack.Pop().(*ast.StringLiteral)
	parent := s.Stack.Peek().(*ast.Constant)
	parent.AcceptStringLiteral(strLit)
}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *SymonAgentListener) EnterIntegerLiteral(ctx *Symonlang.IntegerLiteralContext) {
	integerValue,err := strconv.Atoi(ctx.GetText())
	if err != nil {
		log.Fatalf("string is not an integer literal %s", ctx.GetText())
	}
	sLiteral := &ast.IntegerLiteral{
		Value: int64(integerValue),
	}
	s.Stack.Push(sLiteral)
}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *SymonAgentListener) ExitIntegerLiteral(ctx *Symonlang.IntegerLiteralContext) {
	intLit := s.Stack.Pop().(*ast.IntegerLiteral)
	parent := s.Stack.Peek().(*ast.Constant)
	parent.AcceptIntegerLiteral(intLit)
}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *SymonAgentListener) EnterFloatLiteral(ctx *Symonlang.FloatLiteralContext) {
	floatValue, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		log.Fatalf("string is not an float literal %s", ctx.GetText())
	}
	sLiteral := &ast.FloatLiteral{
		Value: floatValue,
	}
	s.Stack.Push(sLiteral)
}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *SymonAgentListener) ExitFloatLiteral(ctx *Symonlang.FloatLiteralContext) {
	floatLit := s.Stack.Pop().(*ast.FloatLiteral)
	parent := s.Stack.Peek().(*ast.Constant)
	parent.AcceptFloatLiteral(floatLit)
}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *SymonAgentListener) EnterBoolLiteral(ctx *Symonlang.BoolLiteralContext) {
	boolVal, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		log.Fatalf("string is not an boolean literal %s", ctx.GetText())
	}
	sLiteral := &ast.BooleanLiteral{
		Value: boolVal,
	}
	s.Stack.Push(sLiteral)
}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *SymonAgentListener) ExitBoolLiteral(ctx *Symonlang.BoolLiteralContext) {
	boolLit := s.Stack.Pop().(*ast.BooleanLiteral)
	parent := s.Stack.Peek().(*ast.Constant)
	parent.AcceptBooleanLiteral(boolLit)
}