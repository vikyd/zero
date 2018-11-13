package zero

import "reflect"

// IsZeroVal check if any type is its zero value
func IsZeroVal(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// IsDefaultVal alias of IsZeroVal
func IsDefaultVal(x interface{}) bool {
	return IsZeroVal(x)
}
