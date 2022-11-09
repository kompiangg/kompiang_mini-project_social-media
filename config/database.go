package config

import (
	"fmt"
	"log"
	"time"

	model "github.com/kompiang_mini-project_social-media/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Username     string
	Password     string
	Hostname     string
	Port         string
	DatabaseName string
}

var db *gorm.DB

func initDatabase(conf Database) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Hostname,
		conf.Port,
		conf.DatabaseName,
	)

	var err error
	for {
		if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
			break
		}
		log.Println("[HANDLER ERROR] Cant establish the database connection, trying in 1 second")
		time.Sleep(1 * time.Second)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("[HANDLER ERROR] Error getting generic interface object")
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("[INFO] Successfully establishing database connection")
	return nil
}

func migrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.UserDetail{},
		&model.UserSetting{},
		&model.RefreshToken{},
		&model.ChatRoom{},
		&model.ChatRoomUser{},
		&model.Message{},
		&model.Notification{},
		&model.UserRelation{},
		&model.Post{},
		&model.Comment{},
	)
	if err != nil {
		log.Println("[ERROR] Error while create the table")
		return err
	}

	return nil
}

func GetDatabaseConn(conf Database) *gorm.DB {
	if db == nil {
		initDatabase(conf)
		migrateDatabase(db)
	}
	return db
}

func CloseDatabaseConnection() error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("[HANDLER ERROR] Error getting generic interface object")
		return err
	}
	sqlDB.Close()
	return nil
}
