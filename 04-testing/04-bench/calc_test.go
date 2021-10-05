package calc

import "testing"

func TestSum(t *testing.T) {
	result := Sum(5, 7)
	expected := 12

	if result != expected {
		t.Error("Expected:", expected, "Got:", result)
	}
}

// To run, go test -bench=Sum
func BenchmarkSum(b *testing.B){
	for i :=0; i < b.N ; i++{
		Sum(4, 6)
	}
}