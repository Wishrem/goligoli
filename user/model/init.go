package model

import (
	"time"

	"github.com/wishrem/goligoli/pkg/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	dsn := conf.App.UserService.MySQL.Dsn
	_db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger:      logger.Default,
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "goli_",
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := _db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&User{},
		&Role{},
		&Ban{},
	)
	if err != nil {
		panic(err)
	}

	db = _db
}
