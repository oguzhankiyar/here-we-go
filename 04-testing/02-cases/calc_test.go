package calc

import "testing"

func TestSum(t *testing.T) {
	type testCase struct { a, b, sum int }

	cases := []testCase{
		{5, 7, 12},
		{7, 12, 19},
		{3, 5, 8},
	}

	for _, c := range cases {
		result := Sum(c.a, c.b)
		expected := c.sum

		if result != expected {
			t.Error("Expected:", expected, "Got:", result)
		}
	}
}