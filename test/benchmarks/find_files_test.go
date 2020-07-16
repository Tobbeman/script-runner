package benchmarks

import (
	"testing"
)

func BenchmarkHomeBaked(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = homeBaked("test/")
	}
}

func BenchmarkWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = walk("test/")
	}
}
