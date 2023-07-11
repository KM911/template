package main

import (
	"demo/model"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	// 首先插入一些数据
	model.InitDao()
	for i := 0; i < 3000; i++ {
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		model.DC.Create(&newHello)
	}
}

//func BenchmarkListFind(b *testing.B) {
//	model.InitDao()
//	for i := 0; i < b.N; i++ {
//		var hellolist []model.Hello
//		model.DB.Where("age = ?", 180).Find(&hellolist)
//		if len(hellolist) == 0 {
//			// login
//		}
//	}
//}

//func BenchmarkListFirst(b *testing.B) {
//	model.InitDao()
//	for i := 0; i < b.N; i++ {
//		var hellolist []model.Hello
//		model.DB.Where("age = ?", 180).First(&hellolist)
//		if len(hellolist) == 0 {
//			// login
//		}
//	}
//}
//
//// find better then first
//// I just do not know just need to do this
//func BenchmarkHelloFirst(b *testing.B) {
//	model.InitDao()
//	for i := 0; i < b.N; i++ {
//		var Hello model.Hello
//		model.DB.Where("age = ?", 180).First(&Hello)
//		if Hello.ID == 0 {
//
//		}
//	}
//}
//func BenchmarkHelloFind(b *testing.B) {
//	model.InitDao()
//	for i := 0; i < b.N; i++ {
//		var Hello model.Hello
//		model.DB.Where("age = ?", 180).Find(&Hello)
//		if Hello.ID == 0 {
//
//		}
//	}
//}

// 单线程下 利用数据库池 和

type Message struct{}

func BenchmarkSingle(b *testing.B) {

	model.InitDao()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		model.DC.Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		model.DC.Create(&newHell)
	}
}
func BenchmarkDS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		model.DS["db"].Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		model.DS["dc"].Create(&newHell)
	}
}
func BenchmarkSinglego(b *testing.B) {

	model.InitDao()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		go model.DC.Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		go model.DC.Create(&newHell)
	}
}
func BenchmarkDSgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		go model.DS["db"].Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		go model.DS["dc"].Create(&newHell)
	}
}

//BenchmarkList
//BenchmarkList-12           20204             58477 ns/op
//BenchmarkHello
//BenchmarkHello-12           1357            859445 ns/op

// 但是我们这里可能是因为需要print的问题
