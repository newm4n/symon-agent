package ast

type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) Evaluate() (interface{}, error) {
	return fl.Value, nil
}
