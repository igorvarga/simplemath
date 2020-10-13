package simplemath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(2, 2)

	if r != 4 {
		t.Errorf("Add(2, 2) = %f, want 4", r)
	}
}

func TestDivide(t *testing.T) {
	r := Divide(10, 5)

	if r != 2 {
		t.Errorf("Divide(10, 5) = %f, want 2", r)
	}
}

func TestMultiply(t *testing.T) {
	r := Multiply(2, 3)

	if r != 6 {
		t.Errorf("Multiply(2, 3) = %f, want 6", r)
	}
}

func TestSubtract(t *testing.T) {
	r := Subtract(2, 3)

	if r != -1 {
		t.Errorf("Subtract(2, 3) = %f, want -1", r)
	}
}
