package ops

import "testing"

func TestAdd(t *testing.T) {
	if Add(3, 2) != 5 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	if Mul(3, 2) != 6 {
		t.Fail()
	}
}
