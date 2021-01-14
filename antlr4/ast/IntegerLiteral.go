package ast

type IntegerLiteral struct {
	value int64
}

func (il *IntegerLiteral) Evaluate() (interface{}, error) {
	return il.value, nil
}