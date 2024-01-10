package main

import "path/filepath"

func main() {
	ext := filepath.Ext("Makefile")
	println(ext)

}
