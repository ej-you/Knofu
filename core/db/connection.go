package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	userModels "github.com/ej-you/Knofu/user/models"

	"github.com/ej-you/Knofu/settings"
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


// создание таблиц в БД по структурам в Go
func Migrate() {
	settings.InfoLog.Println("Start migrations...")

	// подключение к БД
	db, err := GetConnection()
	settings.DieIf(err)
	// создание таблиц
	createOrRecreateIfExists(db, &userModels.User{})

	settings.InfoLog.Println("Migrated successfully!")
}

// получение соединения с БД
func GetConnection() (*gorm.DB, error) {
	var connection *gorm.DB
	var err error

	connection, err = gorm.Open(mysql.Open(settings.DSN), &gorm.Config{})
	if err != nil {
		return connection, err
	}

	return connection, nil
}
