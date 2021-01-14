package ast

import (
	"fmt"
	"reflect"
	"strings"
)

type Function struct {
	Name string
	Args *ArgumentList
}

func (f *Function) Evaluate() (interface{}, error) {
	argumentValues, err := f.Args.Evaluate()
	if err != nil {
		return nil, fmt.Errorf("function %s got %s", f.Name, err.Error())
	}
	switch f.Name {
	case "ToUpper":
		if argumentValues != nil && len(argumentValues) == 1 {
			if reflect.ValueOf(argumentValues[0]).Kind() == reflect.String {
				return strings.ToUpper(argumentValues[0].(string)), nil
			}
			return nil, fmt.Errorf("function %s requires 1 string argument", f.Name)
		}
		return nil, fmt.Errorf("function %s requires 1 argument", f.Name)
	case "ToLower":
		if argumentValues != nil && len(argumentValues) == 1 {
			if reflect.ValueOf(argumentValues[0]).Kind() == reflect.String {
				return strings.ToUpper(argumentValues[0].(string)), nil
			}
			return nil, fmt.Errorf("function %s requires 1 string argument", f.Name)
		}
		return nil, fmt.Errorf("function %s requires 1 argument", f.Name)
	case "Left":
		if argumentValues != nil && len(argumentValues) == 2 {
			if reflect.ValueOf(argumentValues[0]).Kind() == reflect.String &&
				reflect.ValueOf(argumentValues[1]).Kind() == reflect.Int64 {
				if argumentValues[1].(int64) < 0 {
					return nil, fmt.Errorf("function %s second argument must be positive. Found %d", f.Name, argumentValues[1].(int64))
				}
				if len(argumentValues[0].(string)) < int(argumentValues[1].(int64)) {
					return argumentValues[0].(string), nil
				}
				return argumentValues[0].(string)[0:argumentValues[1].(int64)], nil
			}
			return nil, fmt.Errorf("function %s requires (string, integer) argument", f.Name)
		}
		return nil, fmt.Errorf("function %s requires 2 argument", f.Name)
	case "Length":
		if argumentValues != nil && len(argumentValues) == 1 {
			if reflect.ValueOf(argumentValues[0]).Kind() == reflect.String {
				return int64(len(argumentValues[0].(string))), nil
			}
			return nil, fmt.Errorf("function %s requires 1 string argument", f.Name)
		}
		return nil, fmt.Errorf("function %s requires 1 argument", f.Name)
	default:
		return nil, fmt.Errorf("function %s is not implemented", f.Name)
	}
}
