package prose

import (
	"fmt"
	"testing"
)

// Function name should start with test.
// Name after the test can we whathever I want, need to start capitalized.
// Test functions should accept a single parameter.
func TestTwoElements(t *testing.T) { // Function will be passed a pointer to a
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("No match for expected value") //		testing value.
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if val := JoinWithCommas(list); val != "apple, orange, and pear" {
		fmt.Println(val)
		// Call a method on testing T to fail the test.
		t.Error("No match for expected value") //		testing value.
	}
}
