// 这里是我们使用sqlite的部分

package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path/filepath"
)

var (
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
	hello := connectDatabase(filepath.Join("data", "choice.db"))
	hello.AutoMigrate(&Hello{})
	DS["choice"] = hello
}
