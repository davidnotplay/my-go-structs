package structs

import "testing"

func insertValues(num int) Avl {
	avl := NewAvl()
	// var v Value

	for i := 0; i < num; i++ {
		v := &IntValue{i}
		avl.Insert(v)
	}

	return avl
}

func insertValuesBenchmark(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertValues(num)
	}
}

func Benchmark_AVLInsert1(b *testing.B) {
	insertValuesBenchmark(1, b)
}

func Benchmark_AVLInsert10(b *testing.B) {
	insertValuesBenchmark(10, b)
}

func Benchmark_AVLInsert100(b *testing.B) {
	insertValuesBenchmark(100, b)
}

func Benchmark_AVLInsert1000(b *testing.B) {
	insertValuesBenchmark(1000, b)
}
func Benchmark_AVLInsert10000(b *testing.B) {
	insertValuesBenchmark(10000, b)
}

func Benchmark_AVLInsert100000(b *testing.B) {
	insertValuesBenchmark(100000, b)
}

func Benchmark_AVLInsert1000000(b *testing.B) {
	insertValuesBenchmark(1000000, b)
}

func Benchmark_AVLInsert10000000(b *testing.B) {
	insertValuesBenchmark(1000000, b)
}
