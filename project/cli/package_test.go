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

func Benchmark_Main(b *testing.B) {
	StartLog()
	for i := 0; i < b.N; i++ {
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			log.Println(pos(i), neg(-2*i))
		}
	}
}

func Benchmark_Simple(b *testing.B) {
	StartLog()
	for i := 0; i < b.N; i++ {
		pos, neg := 0, 0
		for i := 0; i < 10; i++ {
			pos += i
			neg += -2 * i
			log.Println(pos, neg)
		}
	}
}
