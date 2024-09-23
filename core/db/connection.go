package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	userModels "github.com/Danil-114195722/Knofu/user/models"

	"github.com/Danil-114195722/Knofu/settings"
)


func createOrRecreateIfExists(db *gorm.DB, table interface{}) {
	var err error

	// если таблица есть, удаляем её
	if db.Migrator().HasTable(table) {
		err = db.Migrator().DropTable(table)
		settings.DieIf(err)
	}
	// создаём таблицу
	err = db.Migrator().CreateTable(table)
	settings.DieIf(err)
}


// создание таблиц в БД по структурам в GO
func Migrate() {
	settings.InfoLog.Println("Start migrations...")
	db := GetConnection()

	createOrRecreateIfExists(db, &userModels.User{})

	settings.InfoLog.Println("Migrated successfully!")
}

// получение соединения с БД
func GetConnection() (db *gorm.DB) {
	settings.InfoLog.Println("Start connection to DB...")

	connection, err := gorm.Open(mysql.Open(settings.DSN), &gorm.Config{})
	settings.DieIf(err)

	settings.InfoLog.Println("Connected to DB successfully!")
	return connection
}
