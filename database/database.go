package database

import (
	"fmt"
	"github.com/restBoard/api/posting/domain"
	"github.com/restBoard/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var setting, _ = config.LoadSetting()
var dsn = fmt.Sprintf("host=%s user=%s password=%s database.Dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
	setting.Database.Host,
	setting.Database.User,
	setting.Database.Password,
	setting.Database.Dbname,
	setting.Database.Port,
)
var Db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func init() {
	Db.AutoMigrate(&domain.Posting{})
	fmt.Println("database.Db Migrated")
}
