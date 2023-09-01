package main

import (
	"demo/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	// 首先插入一些数据
	InitDao()
	for i := 0; i < 3000; i++ {
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		DC.Create(&newHello)
	}
}

//func BenchmarkListFind(b *testing.B) {
//	InitDao()
//	for i := 0; i < b.N; i++ {
//		var hellolist []model.Hello
//		model.DB.Where("age = ?", 180).Find(&hellolist)
//		if len(hellolist) == 0 {
//			// login
//		}
//	}
//}

//func BenchmarkListFirst(b *testing.B) {
//	InitDao()
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
//	InitDao()
//	for i := 0; i < b.N; i++ {
//		var Hello model.Hello
//		model.DB.Where("age = ?", 180).First(&Hello)
//		if Hello.ID == 0 {
//
//		}
//	}
//}
//func BenchmarkHelloFind(b *testing.B) {
//	InitDao()
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

var (
	DB *gorm.DB
	DC *gorm.DB
	DS = make(map[string]*gorm.DB)
)

func connectDatabase(filepath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("failed to connect database", filepath)
		panic(err.Error())
	}
	return db
}

func InitDao() {
	DB = connectDatabase("db.db")
	DC = connectDatabase("dc.db")
	DB.AutoMigrate(&model.Hello{})
	DC.AutoMigrate(&model.Hello{})
	DS["db"] = DB
	DS["dc"] = DC

}

func BenchmarkSingleWriter(b *testing.B) {

	InitDao()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		DC.Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		DC.Create(&newHell)
	}
}
func BenchmarkDoubleWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		DS["db"].Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		DS["dc"].Create(&newHell)
	}
}
func BenchmarkSingleWriterWithGO(b *testing.B) {
	InitDao()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		go DC.Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		go DC.Create(&newHell)
	}
}
func BenchmarkDoubleWriterWithGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond * 1)
		newHello := model.Hello{
			Name: "hello world",
			Age:  18,
		}
		go DS["db"].Create(&newHello)
		newHell := model.Hello{
			Name: "hello world",
			Age:  19,
		}
		go DS["dc"].Create(&newHell)
	}
}
