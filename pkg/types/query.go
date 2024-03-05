package types

import (
	"fmt"
	"reflect"
	"strings"
)

// OrderByDirection type to define our enum for OrderBy direction
type OrderByDirection string

const (
	ASC  OrderByDirection = "ASC"
	DESC OrderByDirection = "DESC"
)

// Operator type to define our enum for supported operators
type Operator string

const (
	Equal         Operator = "="
	NotEqual      Operator = "<>"
	GreaterThan   Operator = ">"
	LessThan      Operator = "<"
	GreaterEquals Operator = ">="
	LessEquals    Operator = "<="
	Like          Operator = "LIKE"
	In            Operator = "IN"
)

// Condition generates a comparison string for a field, given an operator and a value
func Condition(field string, operator Operator, value interface{}) string {
	var formattedValue string

	switch v := value.(type) {
	case string:
		formattedValue = fmt.Sprintf("'%s'", EscapeStringValue(v))
	case int, float64, float32, int64, int32, int16, int8:
		formattedValue = fmt.Sprintf("%v", v)
	case bool:
		formattedValue = fmt.Sprintf("%v", v)
	case []string:
		escapedValues := make([]string, len(v))
		for i, s := range v {
			escapedValues[i] = fmt.Sprintf("'%s'", EscapeStringValue(s))
		}
		formattedValue = fmt.Sprintf("(%s)", strings.Join(escapedValues, ", "))
	case []int, []float64, []float32, []int64, []int32, []int16, []int8:
		// Reflection to handle all int slice types generically
		vals := reflect.ValueOf(v)
		strVals := make([]string, vals.Len())
		for i := 0; i < vals.Len(); i++ {
			strVals[i] = fmt.Sprintf("%d", vals.Index(i).Int())
		}
		formattedValue = fmt.Sprintf("(%s)", strings.Join(strVals, ", "))
	default:
		// Fallback for types not explicitly handled
		formattedValue = fmt.Sprintf("'%v'", EscapeStringValue(fmt.Sprint(v)))
	}

	return fmt.Sprintf("%s %s %s", field, operator, formattedValue)
}

// EscapeStringValue escapes apostrophes in the string by prefixing them with a backslash
func EscapeStringValue(value string) string {
	return strings.ReplaceAll(value, "'", "\\'")
}
