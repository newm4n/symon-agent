package ast

type StringLiteral struct {
	value string
}

func (sl *StringLiteral) Evaluate() (interface{}, error) {
	return sl.value, nil
}
