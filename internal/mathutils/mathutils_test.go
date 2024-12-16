package mathutils

import (
	"testing"
)

func TestGCD(t *testing.T) {
	expect := 9
	result := GCD(18, 189)

	if expect != result {
		t.Fatal()
	}
}
