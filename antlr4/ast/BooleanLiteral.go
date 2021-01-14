package ast

type BooleanLiteral struct {
	value bool
}

func (bl *BooleanLiteral) Evaluate() (interface{}, error){
	return bl.value, nil
}
