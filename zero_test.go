package zero

import (
	"testing"
)

func TestNil(t *testing.T) {
	if !IsZeroVal(nil) {
		t.Error("")
	}
}
func TestInt(t *testing.T) {
	var v int
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestUint(t *testing.T) {
	var v uint
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestFloat64(t *testing.T) {
	var v float64
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestString(t *testing.T) {
	var v string
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestBool(t *testing.T) {
	var v bool
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestStruct(t *testing.T) {
	type T1 struct {
		F1 string
	}
	var v T1
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestChannel(t *testing.T) {
	var v chan int
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestSlice(t *testing.T) {
	var v []int
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestMap(t *testing.T) {
	var v map[string]string
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestByte(t *testing.T) {
	var v byte
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestInterface(t *testing.T) {
	var v interface{}
	if !IsZeroVal(v) {
		t.Error("")
	}
}

func TestPointer(t *testing.T) {
	var v *int
	if !IsZeroVal(v) {
		t.Error("")
	}
}
