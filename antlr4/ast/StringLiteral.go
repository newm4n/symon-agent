package ast

type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) Evaluate() (interface{}, error) {
	return sl.Value, nil
}
