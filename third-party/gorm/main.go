package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TimeTest struct {
	ID int64
	Ts time.Time
	Dt time.Time
}

func (t *TimeTest) TableName() string {
	return "timetest"
}

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/localhost?charset=utf8mb4&parseTime=True&loc=US%2FPacific"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	os.Setenv("TZ", "Africa/Abidjan")
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	var tt TimeTest
	if db.Where("id=?", 1).First(&tt).Error != nil {
		panic("failed to query database")
		return
	}
	fmt.Println(tt)
}
