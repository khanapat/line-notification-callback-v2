package common

import (
	"fmt"
	"testing"
)

func TestFindRuneIndex(t *testing.T) {
	testcases := []struct {
		GivenString string
		GivenRune   rune
		Want        []int
	}{
		{
			GivenString: "bobolol eobo",
			GivenRune:   'o',
			Want:        []int{1, 3, 5, 9, 11},
		},
		{
			GivenString: "trustlnwza555",
			GivenRune:   '.',
			Want:        []int{},
		},
		{
			GivenString: "$trust$lnwza555$",
			GivenRune:   '$',
			Want:        []int{0, 6, 15},
		},
	}

	for _, scenario := range testcases {
		t.Run(fmt.Sprintf("given string:%v rune:%v returns %v", scenario.GivenString, scenario.GivenRune, scenario.Want), func(t *testing.T) {
			givenString := scenario.GivenString
			givenRune := scenario.GivenRune
			want := scenario.Want

			get := FindRuneIndex(givenString, givenRune)
			if !Equal(want, get) {
				t.Errorf("given string:%v rune:%v want %v but got %v\n", scenario.GivenString, scenario.GivenRune, want, get)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
