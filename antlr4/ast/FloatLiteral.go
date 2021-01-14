package ast

type FloatLiteral struct {
	value float64
}

func (fl *FloatLiteral) Evaluate() (interface{}, error) {
	return fl.value, nil
}
