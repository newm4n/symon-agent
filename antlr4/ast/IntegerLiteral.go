package ast

type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) Evaluate() (interface{}, error) {
	return il.Value, nil
}