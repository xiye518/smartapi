package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"smartapi/internal/common"
	"smartapi/internal/log"
)

var DB *gorm.DB

func InitDB(cfg *common.MysqlConfig) error {
	var err error
	//DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		log.Errorf("mysql connect error %s", err)
		return err
	}

	if DB.Error != nil {
		log.Errorf("database error %v", DB.Error)
		return err
	}

	DB.AutoMigrate(&User{})
	return nil
}
