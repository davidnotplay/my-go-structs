package structs

import "testing"

func insertItems(num int) tree {
	tree := tree{}

	for i := 0; i < num; i++ {
		tree.Insert(It(i))
	}

	return tree
}

func insertItemsBenchmark(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertItems(num)
	}
}

func Benchmark_AVLInsert1(b *testing.B) {
	insertItemsBenchmark(1, b)
}

func Benchmark_AVLInsert10(b *testing.B) {
	insertItemsBenchmark(10, b)
}

func Benchmark_AVLInsert100(b *testing.B) {
	insertItemsBenchmark(100, b)
}

func Benchmark_AVLInsert1000(b *testing.B) {
	insertItemsBenchmark(1000, b)
}
func Benchmark_AVLInsert10000(b *testing.B) {
	insertItemsBenchmark(10000, b)
}

func Benchmark_AVLInsert100000(b *testing.B) {
	insertItemsBenchmark(100000, b)
}

func Benchmark_AVLInsert1000000(b *testing.B) {
	insertItemsBenchmark(1000000, b)
}

func Benchmark_AVLInsert10000000(b *testing.B) {
	insertItemsBenchmark(1000000, b)
}
