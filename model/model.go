package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"online_backend/package/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.DBUser,
		setting.DatabaseSetting.DBPassword,
		setting.DatabaseSetting.DBHost,
		setting.DatabaseSetting.DBName))

	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	} else {
		fmt.Println("数据库连接成功.")
	}

	db.AutoMigrate(&User{})

}

func CloseDB() {
	defer db.Close()
}
