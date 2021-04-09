package prose

import (
	"fmt"
	"testing"
)

// Function name should start with test.
// Name after the test can we whathever I want, need to start capitalized.
// Test functions should accept a single parameter.

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			f := format(test.list, test.want, got)
			t.Error(f)
		}
		fmt.Println("Got", got)
	}
}

func TestTwoElements(t *testing.T) { // Function will be passed a pointer to a
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(format(list, want, got)) //		testing value.
	}
}
func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		// Call a method on testing T to fail the test.
		t.Error(format(list, want, got)) //		testing value.
	}
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		// Call a method on testing T to fail the test.
		t.Error(format(list, want, got)) //		testing value.
	}
}

func format(list []string, want string, got string) string {
	var r string
	r += fmt.Sprintln("Input:", list)
	r += fmt.Sprintln("Got  -->", got, "\nWant -->", want)
	return r
}
