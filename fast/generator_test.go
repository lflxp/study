package fast

import (
	"testing"
)

func TestRand_generator_1(t *testing.T) {
	t.Log(Rand_generator_1())
}

func BenchmarkRand_generator_1(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.Log(Rand_generator_1())
	}
}

func TestRand_generator_2(t *testing.T) {
	t.Log(<-Rand_generator_2())
}

func BenchmarkRand_generator_2(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.Log(<-Rand_generator_2())
	}
}

func TestRand_generator_3(t *testing.T) {
	t.Log(<-Rand_generator_3())
}

func BenchmarkRand_generator_3(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.Log(<-Rand_generator_3())
	}
}