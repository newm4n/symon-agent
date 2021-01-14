// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/symon-agen\symonlang.g4 by ANTLR 4.8. DO NOT EDIT.

package Symonlang // symonlang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesymonlangListener is a complete listener for a parse tree produced by symonlangParser.
type BasesymonlangListener struct{}

var _ symonlangListener = &BasesymonlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesymonlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesymonlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesymonlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesymonlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasesymonlangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasesymonlangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction is called when production function is entered.
func (s *BasesymonlangListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasesymonlangListener) ExitFunction(ctx *FunctionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasesymonlangListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasesymonlangListener) ExitConstant(ctx *ConstantContext) {}

// EnterOperator is called when production operator is entered.
func (s *BasesymonlangListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BasesymonlangListener) ExitOperator(ctx *OperatorContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BasesymonlangListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BasesymonlangListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasesymonlangListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasesymonlangListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasesymonlangListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasesymonlangListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BasesymonlangListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BasesymonlangListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BasesymonlangListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BasesymonlangListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}
