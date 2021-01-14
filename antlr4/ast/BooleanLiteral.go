package ast

type BooleanLiteral struct {
	Value bool
}

func (bl *BooleanLiteral) Evaluate() (interface{}, error){
	return bl.Value, nil
}
