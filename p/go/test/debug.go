//go:build debug

package test

import "fmt"

func Breakpoint(a any) {
	fmt.Println(a)
}
