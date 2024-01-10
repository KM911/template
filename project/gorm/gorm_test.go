package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

// 这里我们测试更多的 从 1-10 20 个数据库之间的差异

var (
	DB1  *gorm.DB
	DB2  *gorm.DB
	DB3  *gorm.DB
	DB4  *gorm.DB
	DB5  *gorm.DB
	DB6  *gorm.DB
	DB7  *gorm.DB
	DB8  *gorm.DB
	DB9  *gorm.DB
	DB10 *gorm.DB
	DB11 *gorm.DB
	DB12 *gorm.DB
	DB13 *gorm.DB
	DB14 *gorm.DB
	DB15 *gorm.DB
	DB16 *gorm.DB
	DB17 *gorm.DB
	DB18 *gorm.DB
	DB19 *gorm.DB
	DB20 *gorm.DB
	DS   = make(map[string]*gorm.DB)
)

type Message struct {
	ID   int64 `gorm:"primary_key"`
	Msg  string
	Date string
}

func connect(filename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func init() {
	// 好无聊啊 就是说 其实是没有什么感觉的 
}

// 这里的影响因素是已经写入的数据是否会对插入有性能影响
func BenchmarkDB_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 20; i++ {
			newMessage := Message{
				Msg:  "hello world",
				Date: "2021-10-10",
			}
			DB1.Create(&newMessage)
		}
	}
}

func BenchmarkDB_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			newMessage := Message{
				Msg:  "hello world",
				Date: "2021-10-10",
			}
			DB1.Create(&newMessage)
			DB2.Create(&newMessage)
		}
	}
}

func BenchmarkDB_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 4; i++ {
			newMessage := Message{
				Msg:  "hello world",
				Date: "2021-10-10",
			}
			DB1.Create(&newMessage)
			DB2.Create(&newMessage)
			DB3.Create(&newMessage)
			DB4.Create(&newMessage)
			DB5.Create(&newMessage)

		}

	}

}





