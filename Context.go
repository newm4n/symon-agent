package symon_agen

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type MonitorContext map[string]string

func NewMonitorContext() MonitorContext {
	return make(MonitorContext)
}

func(ctx MonitorContext) String() string {
	var buffer bytes.Buffer
	keys := make([]string, 0)
	for k := range ctx {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := range keys {
		buffer.WriteString(fmt.Sprintf("%s : %s\n", keys[i],ctx[keys[i]]))
	}
	return buffer.String()
}

func (ctx MonitorContext) GetInt(key string) int {
	intValue, err := strconv.Atoi(ctx[key])
	if err != nil {
		return 0
	}
	return intValue
}
func (ctx MonitorContext) GetBool(key string) bool {
	switch strings.ToUpper(ctx[key]) {
	case "TRUE":
		return true
	case "YES":
		return true
	default:
		return false
	}
}
func (ctx MonitorContext) GetFloat(key string) float64 {
	floatValue, err := strconv.ParseFloat(ctx[key], 64)
	if err != nil {
		return 0
	}
	return floatValue
}

func (ctx MonitorContext) SetInt(key string, value int64) {
	ctx[key] = fmt.Sprintf("%d", value)
}
func (ctx MonitorContext) SetBool(key string, value bool) {
	ctx[key] = fmt.Sprintf("%v", value)
}
func (ctx MonitorContext) SetFloat(key string, value float64) {
	ctx[key] = fmt.Sprintf("%f", value)
}
func (ctx MonitorContext) Contains(key string) bool {
	_, exist := ctx[key]
	return exist
}

