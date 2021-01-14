package ast

type Function struct {
	name string
	args *ArgumentList
}

func (e *Function) Evaluate() (interface{}, error) {
	return nil, nil
}