package cli

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func StartLog() {
	file := "log.txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 可以作为引用计数器
func Test_Main(b *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

//func Benchmark_Main(b *testing.B) {
//	StartLog()
//	for i := 0; i < b.N; i++ {
//		pos, neg := adder(), adder()
//		for i := 0; i < 10; i++ {
//			log.Println(pos(i), neg(-2*i))
//		}
//	}
//}
//
//func Benchmark_Simple(b *testing.B) {
//	StartLog()
//	for i := 0; i < b.N; i++ {
//		pos, neg := 0, 0
//		for i := 0; i < 10; i++ {
//			pos += i
//			neg += -2 * i
//			log.Println(pos, neg)
//		}
//	}
//}

//func BenchmarkPrintln(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		println(i)
//	}
//}
//
//func BenchmarkPrint(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		print(i)
//	}
//}
//
//func BenchmarkPrintf(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		fmt.Printf("%d", i)
//	}
//}
//
//func BenchmarkFmtPrint(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		fmt.Println(i)
//	}
//}

// 使用日志可以大幅度提升我们的程序性能 当然了 我们还要做一个

func MatrixSum() int {
	martix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		martix[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			martix[i][j] = i + j
		}
	}
	sum := 0
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			sum += martix[row][col]
		}
	}
	return sum
}
func BenchmarkName(b *testing.B) {
	StartSimpleLog("withone.log")

	for i := 0; i < b.N; i++ {
		//MatrixSum()
		log.Println(MatrixSum())
	}
}

func StartSimpleLog(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func BenchmarkWithoutLog(b *testing.B) {
	//StartLog()
	StartSimpleLog("without.log")
	for i := 0; i < b.N; i++ {
		log.Println(MatrixSum())
		log.Print(i)
	}
}

//BenchmarkPrintln-12                27380             43433 ns/op
//BhmarkPrintf-12               37029             34927 ns/op
//BenchmarkFmtPrint
//BenchmarkFmtPrint-12               31683             35794 ns/op
//BenchmarkName
//BenchmarkName-12                  585468              2069 ns/op
