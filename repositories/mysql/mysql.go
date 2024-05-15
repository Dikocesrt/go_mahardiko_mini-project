package mysql

import (
	"fmt"
	activityDb "habit/repositories/mysql/activity"
	adminDb "habit/repositories/mysql/admin"
	expertDb "habit/repositories/mysql/expert"
	hireDb "habit/repositories/mysql/hire"
	userDb "habit/repositories/mysql/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBUnix string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	// dsn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", config.DBUser, config.DBPass, config.DBUnix, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userDb.User{}, &activityDb.Activity{}, &activityDb.ActivityType{}, &activityDb.ActivityDetail{}, &expertDb.Expert{}, &expertDb.BankAccount{}, &expertDb.BankAccountType{}, &expertDb.Expertise{}, &hireDb.Hire{}, &adminDb.Admin{})
}