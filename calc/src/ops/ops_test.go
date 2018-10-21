package ops

import "testing"

func ExampleAdd() {
	Add(3, 2)
	// Output:
	// -
}

func ExampleMul() {
	Mul(2, 2)
	Mul(3, 3)
	// Output:
	// -
}

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
