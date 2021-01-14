package ast

import (
	"reflect"
	"testing"
)

func TestExpression_EvaluateAddition(t *testing.T) {
	testData := []*struct {
		left interface{}
		right interface{}
		operator string
		isError bool
		expect interface{}
	}{
		{ left : int64(123), right : int64(123), operator: "+", isError : false, expect : int64(246)},
		{ left : int64(123), right : float64(123), operator: "+", isError : false, expect : float64(246)},
		{ left : int64(123), right : "123", operator: "+", isError : false, expect : "123123"},
		{ left : int64(123), right : true, operator: "+", isError : true, expect: nil},

		{ left : float64(123), right : int64(123), operator: "+", isError : false, expect : float64(246)},
		{ left : float64(123), right : float64(123), operator: "+", isError : false, expect : float64(246)},
		{ left : float64(123), right : "123", operator: "+", isError : false, expect : "123.000000123"},
		{ left : float64(123), right : true, operator: "+", isError : true, expect: nil},

		{ left : "123", right : int64(123), operator: "+", isError : false, expect : "123123"},
		{ left : "123", right : float64(123), operator: "+", isError : false, expect : "123123.000000"},
		{ left : "123", right : "123", operator: "+", isError : false, expect : "123123"},
		{ left : "123", right : true, operator: "+", isError : false, expect: "123true"},

		{ left : false, right : int64(123), operator: "+", isError : true, expect : int64(246)},
		{ left : false, right : float64(123), operator: "+", isError : true, expect : float64(246)},
		{ left : false, right : "123", operator: "+", isError : false, expect : "false123"},
		{ left : false, right : true, operator: "+", isError : true, expect: nil},


		{ left : int64(123), right : int64(23), operator: "-", isError : false, expect : int64(100)},
		// add more sample for
		// - * / % && || == != < <= > >=
	}

	for i, td := range testData {
		var LeftExpr *Expression
		var RightExpr *Expression

		switch reflect.ValueOf(td.left).Kind() {
		case reflect.Int64:
			LeftExpr = &Expression{
				constant:        &Constant{integer: &IntegerLiteral{Value: td.left.(int64)}},
			}
		case reflect.Float64:
			LeftExpr = &Expression{
				constant:        &Constant{float: &FloatLiteral{Value: td.left.(float64)}},
			}
		case reflect.String:
			LeftExpr = &Expression{
				constant:        &Constant{str: &StringLiteral{Value: td.left.(string)}},
			}
		case reflect.Bool:
			LeftExpr = &Expression{
				constant:        &Constant{bool: &BooleanLiteral{Value: td.left.(bool)}},
			}
		default:
			t.Errorf("%d -> Left Not supported kind %s", i, reflect.ValueOf(td.left).Kind().String())
		}

		switch reflect.ValueOf(td.right).Kind() {
		case reflect.Int64:
			RightExpr = &Expression{
				constant:        &Constant{integer: &IntegerLiteral{Value: td.right.(int64)}},
			}
		case reflect.Float64:
			RightExpr = &Expression{
				constant:        &Constant{float: &FloatLiteral{Value: td.right.(float64)}},
			}
		case reflect.String:
			RightExpr = &Expression{
				constant:        &Constant{str: &StringLiteral{Value: td.right.(string)}},
			}
		case reflect.Bool:
			RightExpr = &Expression{
				constant:        &Constant{bool: &BooleanLiteral{Value: td.right.(bool)}},
			}
		default:
			t.Errorf("%d -> Right Not supported kind %s", i, reflect.ValueOf(td.right).Kind().String())
		}

		MainExpression := &Expression{
			leftExpression:  LeftExpr,
			rightExpression: RightExpr,
			operator:        &Operator{Operator: td.operator},
		}

		intf, err := MainExpression.Evaluate()
		if err == nil && td.isError {
			t.Errorf("%d -> Expected to have error but yield no error", i)
		} else if err != nil && !td.isError {
			t.Errorf("%d -> Expected to have no error but yield error", i)
		} else if err == nil && !td.isError {
			if reflect.ValueOf(intf).Kind() != reflect.ValueOf(td.expect).Kind() {
				t.Errorf("%d -> Have different result type expect %s but %s", i, reflect.ValueOf(td.expect).Kind().String(), reflect.ValueOf(intf).Kind().String())
			} else {
				switch reflect.ValueOf(td.expect).Kind() {
					case reflect.Int64:
						if td.expect.(int64) != intf.(int64) {
							t.Errorf("%d -> Expect %d but %d", i, td.expect.(int64), intf.(int64))
						}
					case reflect.Float64:
						if td.expect.(float64) != intf.(float64) {
							t.Errorf("%d -> Expect %f but %f", i, td.expect.(float64), intf.(float64))
						}
					case reflect.String:
						if td.expect.(string) != intf.(string) {
							t.Errorf("%d -> Expect \"%s\" but \"%s\"", i, td.expect.(string), intf.(string))
						}
					case reflect.Bool:
						if td.expect.(bool) != intf.(bool) {
							t.Errorf("%d -> Expect %v but %v", i, td.expect.(bool), intf.(bool))
						}
				}
			}
		}
	}
}


