package antlr4

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	Symonlang "github.com/newm4n/symon-agen/antlr4/symonlang"
)


type SymonAgentListener struct {
	Symonlang.BasesymonlangListener
}

// EnterEveryRule is called when any rule is entered.
func (s *SymonAgentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *SymonAgentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *SymonAgentListener) EnterExpression(ctx *Symonlang.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *SymonAgentListener) ExitExpression(ctx *Symonlang.ExpressionContext) {}

// EnterFunction is called when production function is entered.
func (s *SymonAgentListener) EnterFunction(ctx *Symonlang.FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *SymonAgentListener) ExitFunction(ctx *Symonlang.FunctionContext) {}

// EnterConstant is called when production constant is entered.
func (s *SymonAgentListener) EnterConstant(ctx *Symonlang.ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *SymonAgentListener) ExitConstant(ctx *Symonlang.ConstantContext) {}

// EnterOperator is called when production operator is entered.
func (s *SymonAgentListener) EnterOperator(ctx *Symonlang.OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *SymonAgentListener) ExitOperator(ctx *Symonlang.OperatorContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *SymonAgentListener) EnterArgumentList(ctx *Symonlang.ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *SymonAgentListener) ExitArgumentList(ctx *Symonlang.ArgumentListContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *SymonAgentListener) EnterStringLiteral(ctx *Symonlang.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *SymonAgentListener) ExitStringLiteral(ctx *Symonlang.StringLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *SymonAgentListener) EnterIntegerLiteral(ctx *Symonlang.IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *SymonAgentListener) ExitIntegerLiteral(ctx *Symonlang.IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *SymonAgentListener) EnterFloatLiteral(ctx *Symonlang.FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *SymonAgentListener) ExitFloatLiteral(ctx *Symonlang.FloatLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *SymonAgentListener) EnterBoolLiteral(ctx *Symonlang.BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *SymonAgentListener) ExitBoolLiteral(ctx *Symonlang.BoolLiteralContext) {}