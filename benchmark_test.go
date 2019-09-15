package mygostructs

import "testing"

// List benchmarks
// ---------------
func listInsertItem(num int) *List {
	l := NewList(true)

	for i := 1; i <= num; i++ {
		l.AddAfter(It(i))
	}

	return &l
}

func searchInListNElem(num int, b *testing.B) *List {
	println("prev")
	list := listInsertItem(num)
	println("finish")

	for i := 0; i < b.N; i++ {
		list.Search(It(i))
	}

	return list
}

// func Benchmark_ListSearch100(b *testing.B) {
// 	searchInListNElem(100, b)
// }

// func Benchmark_ListSearch1000(b *testing.B) {
// 	searchInListNElem(1000, b)
// }

// func Benchmark_ListSearch10000(b *testing.B) {
// 	searchInListNElem(10000, b)
// }

// func Benchmark_ListSearch100000(b *testing.B) {
// 	searchInListNElem(100000, b)
// }

// func Benchmark_ListSearch1000000(b *testing.B) {
// 	searchInListNElem(1000000, b)
// }

func Benchmark_ListSearch10000000(b *testing.B) {
	searchInListNElem(10000000, b)
}
