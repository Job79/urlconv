package urlconv

import (
	"net/url"
	"testing"
	"time"
)

type test struct {
	String string    `url:"S"`
	Int    int       `url:"I"`
	Float  float64   `url:"F"`
	Array  []string  `url:"A"`
	Bool   bool      `url:"B"`
	Time   time.Time `url:"T"`
}

// TestUnmarshal tests whether the unmarshal function works correctly
func TestUnmarshal(t *testing.T) {
	values := url.Values{
		"S": []string{"string"},
		"I": []string{"1"},
		"F": []string{"1.1"},
		"A": []string{"a", "b"},
		"B": []string{"true"},
		"T": []string{"2016-01-01T00:00:00Z"},
	}

	var s test
	Unmarshal(values, &s)

	if s.String != "string" {
		t.Errorf("string = %s, want string", s.String)
	} else if s.Int != 1 {
		t.Errorf("Int = %d, want 1", s.Int)
	} else if s.Float != 1.1 {
		t.Errorf("float = %f, want 1.1", s.Float)
	} else if len(s.Array) != 2 {
		t.Errorf("Array = %v, want [a b]", s.Array)
	} else if s.Bool != true {
		t.Errorf("Bool = %v, want [true]", s.Bool)
	} else if s.Time != time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Time = %v, want [2016-01-01T00:00:00Z]", s.Time)
	}
}
