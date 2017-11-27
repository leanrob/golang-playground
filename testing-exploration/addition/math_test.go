package addition

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Log("Failed to add numbers")
		t.Fail()
	}

	result = Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("Cannot add more than 2 numbers")
	}
}
