package fast

import "testing"

func TestGoForFunc(t *testing.T) {
	GoForFunc(100)
}

func BenchmarkGoForFunc(b *testing.B) {
	GoForFunc(b.N)
}