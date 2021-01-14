package ast

import "fmt"

type Constant struct {
	bool *BooleanLiteral
	str *StringLiteral
	integer *IntegerLiteral
	float *FloatLiteral
}

func (c *Constant) Evaluate() (interface{}, error){
	if c.bool != nil {
		return c.bool.Evaluate()
	}
	if c.str != nil {
		return c.str.Evaluate()
	}
	if c.integer != nil {
		return c.integer.Evaluate()
	}
	if c.float != nil {
		return c.float.Evaluate()
	}
	return nil, fmt.Errorf("unknown constant value")
}