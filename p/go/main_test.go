package main

import "testing"

//func BenchmarkWithoutPool(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		buffer := new(bytes.Buffer)
//		buffer.WriteString("Hello, world!")
//	}
//}
//
//func BenchmarkWithPool(b *testing.B) {
//	b.ReportAllocs()
//	pool := sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
//	for i := 0; i < b.N; i++ {
//		buffer := pool.Get().(*bytes.Buffer)
//		buffer.WriteString("Hello, world!")
//		pool.Put(buffer)
//	}
//}

func BenchmarkHello(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

	}
}
func BenchmarkHello2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

	}
}
