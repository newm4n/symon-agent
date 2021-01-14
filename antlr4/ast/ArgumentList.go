package ast

import "fmt"

type ArgumentList struct {
	arguments []*Expression
}

func (al *ArgumentList) Evaluate() ([]interface{}, error) {
	args := make([]interface{},0)
	if al.arguments == nil {
		return args, nil
	}
	for i, v := range al.arguments {
		actual, err := v.Evaluate()
		if err != nil {
			return nil, fmt.Errorf("argument %d got %s", i, err.Error())
		}
		args = append(args, actual)
	}
	return args, nil
}

func (al *ArgumentList) AcceptChildExpression(expr *Expression) {
	if al.arguments == nil {
		al.arguments = make([]*Expression, 0)
	}
	al.arguments = append(al.arguments, expr)
}
