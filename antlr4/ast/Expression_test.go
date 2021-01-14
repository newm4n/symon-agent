package ast

import (
	"reflect"
	"testing"
)

func TestExpression_EvaluateAddition(t *testing.T) {
	LeftExpr := &Expression{
		constant:        &Constant{integer: &IntegerLiteral{value: 123}},
	}
	RightExpr := &Expression{
		constant:        &Constant{integer: &IntegerLiteral{value: 321}},
	}
	MainExpression := &Expression{
		leftExpression:  LeftExpr,
		rightExpression: RightExpr,
		operator:        &Operator{operator: "+"},
	}

	intf, err := MainExpression.Evaluate()
	if err != nil {
		t.Error(err.Error())
	}
	if reflect.ValueOf(intf).Kind() != reflect.Int64 {
		t.Error("Return type is not an int64")
	}
	val := intf.(int64)
	if val != 444 {
		t.Error("Value is not 444 but ", val)
	}
}
