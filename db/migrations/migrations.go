package migrations

import (
	"git.si-media.biz/takutoarao/start/db"
	"git.si-media.biz/takutoarao/start/models"
)

// Execute マイグレーション実行
func Execute() {
	migrate(
		// add some migration struct
		// &models.Categories{},
		&models.ReportCategory{},
	)
}

func migrate(values ...interface{}) {
	for _, value := range values {
		db.ConnectDB().AutoMigrate(value)
	}
}
