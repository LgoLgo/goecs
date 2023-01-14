package migrate

import (
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"srvs/user_srv/model"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/LgoECS_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL Threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color printing
		},
	)

	// global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("user%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}
}
