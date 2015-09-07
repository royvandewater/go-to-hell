package challenge2

import "testing"

func TestCase1(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{120, false},
		{166, true},
		{141, true},
		{79, false},
		{26, true},
		{158, true},
		{174, false},
		{141, true},
		{169, true},
		{129, false},
		{199, false},
		{27, false},
		{57, true},
		{183, true},
		{173, false},
		{5, false},
		{111, true},
		{145, true},
		{59, false},
		{64, false},
	}
	for _, c := range cases {
		got := Run(c.in)
		if got != c.want {
			t.Errorf("Run(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
