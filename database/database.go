package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabaseConnection() {
	/*logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
		)*/
	dsn := "root:1q2w3e4r5t@tcp(127.0.0.1)/qmove_backup"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Unable to connect with database")
	}
	DB = database

}
