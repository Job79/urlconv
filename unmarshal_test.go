package urlconv

import (
	"net/url"
	"testing"
)

type test struct {
	String string   `url:"S"`
	Int    int      `url:"I"`
	Array  []string `url:"A"`
	Bool   bool     `url:"B"`
}

// TestUnmarshal tests whether the unmarshal function works correctly
func TestUnmarshal(t *testing.T) {
	values := url.Values{
		"S": []string{"string"},
		"I": []string{"1"},
		"A": []string{"a", "b"},
		"B": []string{"true"},
	}

	var s test
	Unmarshal(values, &s)

	if s.String != "string" {
		t.Errorf("string = %s, want string", s.String)
	} else if s.Int != 1 {
		t.Errorf("Int = %d, want 1", s.Int)
	} else if len(s.Array) != 2 {
		t.Errorf("Array = %v, want [a b]", s.Array)
	} else if s.Bool != true {
		t.Errorf("Bool = %v, want [true]", s.Bool)
	}
}
