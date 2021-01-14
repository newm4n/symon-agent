package symon_agen

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)


type JobContext interface {
	fmt.Stringer

	Set(key, value string)
	SetInt(key string, value int64)
	SetBool(key string, value bool)
	SetFloat(key string, value float64)

	Contains(key string) bool

	Get(key string) string
	GetMap() map[string]string
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat(key string) float64
}

func NewDefaultContext() JobContext {
	return &DefaultContext{
		data: make(map[string]string),
	}
}

type DefaultContext struct {
	data map[string]string
}

func(ctx *DefaultContext) String() string {
	var buffer bytes.Buffer

	keys := make([]string, 0)
	for k := range ctx.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := range keys {
		buffer.WriteString(fmt.Sprintf("%s : %s\n", keys[i],ctx.data[keys[i]]))
	}

	return buffer.String()
}

func (ctx *DefaultContext) Get(key string) string {
	return ctx.data[key]
}
func (ctx *DefaultContext) GetInt(key string) int {
	intValue, err := strconv.Atoi(ctx.Get(key))
	if err != nil {
		return 0
	}
	return intValue
}
func (ctx *DefaultContext) GetBool(key string) bool {
	switch strings.ToUpper(ctx.Get(key)) {
	case "TRUE":
		return true
	case "YES":
		return true
	default:
		return false
	}
}
func (ctx *DefaultContext) GetFloat(key string) float64 {
	floatValue, err := strconv.ParseFloat(ctx.Get(key), 64)
	if err != nil {
		return 0
	}
	return floatValue
}

func (ctx *DefaultContext) Set(key string, value string) {
	ctx.data[key] = value
}
func (ctx *DefaultContext) SetInt(key string, value int64) {
	ctx.Set(key, fmt.Sprintf("%d", value))
}
func (ctx *DefaultContext) SetBool(key string, value bool) {
	ctx.Set(key, fmt.Sprintf("%v", value))
}
func (ctx *DefaultContext) SetFloat(key string, value float64) {
	ctx.Set(key, fmt.Sprintf("%f", value))
}
func (ctx *DefaultContext) Contains(key string) bool {
	_, exist := ctx.data[key]
	return exist
}
func (ctx *DefaultContext) GetMap() map[string]string {
	return ctx.data
}
