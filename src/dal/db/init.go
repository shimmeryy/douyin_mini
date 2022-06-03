package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tiktok/src/config"
)

var DB *gorm.DB

func Init() {
	var err error
	//连接地址
	dsn := config.AppConfig.DataSource.UserName + ":" + config.AppConfig.DataSource.Password + "@" + config.AppConfig.DataSource.Url
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info), //打印执行的sql
		},
	)
	if err != nil {
		panic(err)
	}

	//初始化表结构
	InitTable()
}

func InitTable() {
	//进行表的创建
	m := DB.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	if !m.HasTable(&Video{}) {
		if err := m.CreateTable(&Video{}); err != nil {
			panic(err)
		}
	}
	if !m.HasTable(&Comment{}) {
		if err := m.CreateTable(&Comment{}); err != nil {
			panic(err)
		}
	}
	if !m.HasTable(&Favor{}) {
		if err := m.CreateTable(&Favor{}); err != nil {
			panic(err)
		}
	}
}
