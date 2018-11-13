package sqrt

import (
	"math"
	"testing"
)

const test = 1012.1234215
const expected = 31.813887

func TestPythonSqrt(t *testing.T) {
	result, err := pythonSqrt(test)

	if result != expected {
		t.Errorf("math.Sqrt(%f) should be %f; got %f", test, expected, result)
	}

	if err != nil {
		t.Error(err)
	}
}

func BenchmarkPythonSqrt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pythonSqrt(test)
	}
}

func BenchmarkGoSqrt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		math.Sqrt(test)
	}
}
