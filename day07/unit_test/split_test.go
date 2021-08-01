package unittest

import (
	"reflect"
	"testing"
)

// 单元测试
func TestSplit(t *testing.T) {
	want := []string{"a", "c", "d"}
	got := Split("abcbde", "b")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// 性能测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("192.168.1.2", ".")
	}
}

func benchMarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchMarkFib(b, 1)
}

func BenchmarkFib10(b *testing.B) {
	benchMarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchMarkFib(b, 20)
}

func BenchmarkFib40(b *testing.B) {
	benchMarkFib(b, 40)
}
