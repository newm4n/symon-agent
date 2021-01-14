package ast

import (
	"fmt"
	"reflect"
	"strings"
)

type Expression struct {
	constant *Constant
	function *Function
	leftExpression *Expression
	rightExpression *Expression
	operator *Operator
}

func (e *Expression) Evaluate() (interface{}, error) {
	if e.constant != nil {
		return e.constant.Evaluate()
	}
	if e.function != nil {
		return e.function.Evaluate()
	}
	if e.leftExpression != nil && e.rightExpression != nil && e.operator != nil  {
		lInterface, err := e.leftExpression.Evaluate()
		if err != nil {
			return nil, err
		}
		rInterface, err := e.rightExpression.Evaluate()
		if err != nil {
			return nil, err
		}
		op := e.operator.operator
		switch op {
		case "+" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return fmt.Sprintf("%s%s", lInterface.(string), rInterface.(string)), nil
				case reflect.Int64:
					return fmt.Sprintf("%s%d", lInterface.(string), rInterface.(int64)), nil
				case reflect.Float64:
					return fmt.Sprintf("%s%f", lInterface.(string), rInterface.(float64)), nil
				case reflect.Bool:
					return fmt.Sprintf("%s%v", lInterface.(string), rInterface.(bool)), nil
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return fmt.Sprintf("%d%s", lInterface.(int64), rInterface.(string)), nil
				case reflect.Int64:
					return lInterface.(int64) + rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) + rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not ADD an integer to a boolean value")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return fmt.Sprintf("%f%s", lInterface.(float64), rInterface.(string)), nil
				case reflect.Int64:
					return lInterface.(float64)-float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) + rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not ADD a float to a boolean value")
				}
			case reflect.Bool:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return fmt.Sprintf("%v%s", lInterface.(bool), rInterface.(string)), nil
				case reflect.Int64:
					return nil, fmt.Errorf("can not ADD a boolean to an integer value")
				case reflect.Float64:
					return nil, fmt.Errorf("can not ADD a boolean to a float value")
				case reflect.Bool:
					return nil, fmt.Errorf("can not ADD between two boolean values")
				}
			}
		case "-" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not SUB a string")
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not SUB a integer with a string")
				case reflect.Int64:
					return lInterface.(int64) - rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) - rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not SUB a integer with a boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not SUB a float with a string")
				case reflect.Int64:
					return lInterface.(float64) - float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) - rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not SUB a float to a boolean value")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not SUB a boolean")
			}
		case "/" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not DIV a string")
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not DIV an integer with string")
				case reflect.Int64:
					return lInterface.(int64) / rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) / rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not DIV an integer with boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not DIV a float with string")
				case reflect.Int64:
					return lInterface.(float64) / float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) / rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not DIV a float with boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not DIV a boolean")
			}
		case "*" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not MUL a string")
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not MUL an integer with string")
				case reflect.Int64:
					return lInterface.(int64) * rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) * rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not MUL an integer with boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not MUL a float with string")
				case reflect.Int64:
					return lInterface.(float64) * float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) * rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not MUL a float with boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not MUL a boolean")
			}
		case "%" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not MOD a string")
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not MOD an integer with string")
				case reflect.Int64:
					return lInterface.(int64) % rInterface.(int64), nil
				case reflect.Float64:
					return nil, fmt.Errorf("can not MOD a float")
				case reflect.Bool:
					return nil, fmt.Errorf("can not MOD an integer with boolean")
				}
			case reflect.Float64:
				return nil, fmt.Errorf("can not MOD a float")
			case reflect.Bool:
				return nil, fmt.Errorf("can not MOD a boolean")
			}
		case "&&" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not LOGIC AND a string")
			case reflect.Int64:
				return nil, fmt.Errorf("can not LOGIC AND an integer")
			case reflect.Float64:
				return nil, fmt.Errorf("can not LOGIC AND a float")
			case reflect.Bool:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Bool:
					return lInterface.(bool) && rInterface.(bool), nil
				default:
					return nil, fmt.Errorf("can not LOGIC AND a boolean to other than boolean")
				}
			}
		case "||" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				return nil, fmt.Errorf("can not LOGIC OR a string")
			case reflect.Int64:
				return nil, fmt.Errorf("can not LOGIC OR an integer")
			case reflect.Float64:
				return nil, fmt.Errorf("can not LOGIC OR a float")
			case reflect.Bool:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Bool:
					return lInterface.(bool) || rInterface.(bool), nil
				default:
					return nil, fmt.Errorf("can not LOGIC OR a boolean to other than boolean")
				}
			}
		case "==" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return lInterface.(string) == rInterface.(string), nil
				default:
					return false, nil
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Int64:
					return lInterface.(int64) == rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) == rInterface.(float64), nil
				default:
					return false, nil
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Int64:
					return lInterface.(float64) == float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) == rInterface.(float64), nil
				default:
					return false, nil
				}
			case reflect.Bool:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Bool:
					return lInterface.(bool) == rInterface.(bool), nil
				default:
					return false, nil
				}
			}
		case "!=" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return lInterface.(string) != rInterface.(string), nil
				default:
					return false, nil
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Int64:
					return lInterface.(int64) != rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) != rInterface.(float64), nil
				default:
					return false, nil
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Int64:
					return lInterface.(float64) != float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) != rInterface.(float64), nil
				default:
					return false, nil
				}
			case reflect.Bool:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.Bool:
					return lInterface.(bool) != rInterface.(bool), nil
				default:
					return false, nil
				}
			}
		case "<" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return strings.Compare(lInterface.(string), rInterface.(string)) < 0, nil
				default:
					return nil, fmt.Errorf("can not LT a string to other than string")
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not LT an integer to string")
				case reflect.Int64:
					return lInterface.(int64) < rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) < rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not LT an integer to boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not LT a float to string")
				case reflect.Int64:
					return lInterface.(float64) < float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) < rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not LT a float to boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not LT a boolean")
			}
		case "<=" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return strings.Compare(lInterface.(string), rInterface.(string)) <= 0, nil
				default:
					return nil, fmt.Errorf("can not LTE a string to other than string")
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not LTE an integer to string")
				case reflect.Int64:
					return lInterface.(int64) <= rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) <= rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not LTE an integer to boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not LTE a float to string")
				case reflect.Int64:
					return lInterface.(float64) <= float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) <= rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not LTE a float to boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not LT a boolean")
			}
		case ">" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return strings.Compare(lInterface.(string), rInterface.(string)) > 0, nil
				default:
					return nil, fmt.Errorf("can not GT a string to other than string")
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not GT an integer to string")
				case reflect.Int64:
					return lInterface.(int64) > rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) > rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not GT an integer to boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not GT a float to string")
				case reflect.Int64:
					return lInterface.(float64) > float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) > rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not GT a float to boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not GT a boolean")
			}
		case ">=" :
			switch reflect.ValueOf(lInterface).Kind() {
			case reflect.String:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return strings.Compare(lInterface.(string), rInterface.(string)) >= 0, nil
				default:
					return nil, fmt.Errorf("can not GTE a string to other than string")
				}
			case reflect.Int64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not GTE an integer to string")
				case reflect.Int64:
					return lInterface.(int64) >= rInterface.(int64), nil
				case reflect.Float64:
					return float64(lInterface.(int64)) >= rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not GTE an integer to boolean")
				}
			case reflect.Float64:
				switch reflect.ValueOf(rInterface).Kind() {
				case reflect.String:
					return nil, fmt.Errorf("can not GTE a float to string")
				case reflect.Int64:
					return lInterface.(float64) >= float64(rInterface.(int64)) , nil
				case reflect.Float64:
					return lInterface.(float64) >= rInterface.(float64), nil
				case reflect.Bool:
					return nil, fmt.Errorf("can not GTE a float to boolean")
				}
			case reflect.Bool:
				return nil, fmt.Errorf("can not GTE a boolean")
			}
		default :
			return nil, fmt.Errorf("not yet implemented")
		}
	}
	return nil, fmt.Errorf("some logic not yet implemented")
}