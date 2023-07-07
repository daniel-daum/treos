package hello

import "testing"

func TestAddition(t *testing.T) {

	result := Addition(2, 2)

	if result != 4 {
		t.Errorf("Result was incorrect, got: %q.", result)
	}

}
