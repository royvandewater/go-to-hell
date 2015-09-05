package challenge1

import "testing"

func TestCase1(t *testing.T) {
  cases := []struct {
		in int
    want uint64
	}{
		{0, 0},
		{1, 1},
		{67, 44945570212853},
		{72, 498454011879264},
		{83, 99194853094755497},
	}
  for _, c := range cases {
		got := Run(c.in)
		if got != c.want {
			t.Errorf("Run(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
