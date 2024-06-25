package connection

import (
	"fmt"
	"log"

	"github.com/robert-self-secret/go-init-project.git/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBOpenConenction(cnf *config.Config) *gorm.DB {
	GormUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.Database.Username, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.DBName)
	fmt.Println(GormUrl)
	dialect := mysql.Open(GormUrl)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed create connection to DB : %s", err.Error())
	}
	return db
}
