package database

import (
	"fmt"
	"httpsink/internal/config"
	"httpsink/internal/database/models"
	"log/slog"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var appDb *gorm.DB // singleton
var createAppDbOnce sync.Once

func GetDatabase() *gorm.DB {
	createAppDbOnce.Do(func() {
		slog.Info("Creating a new database connection")
		var err error
		appDb, err = gorm.Open(sqlite.Open(config.SQLITE_DB), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		err = appDb.AutoMigrate(&models.Room{}, &models.TestCase{}, &models.TestResult{})
		if err != nil {
			slog.Error(fmt.Sprintf("Could not migrate the database %v", err))
			return
		}
	})
	return appDb
}
