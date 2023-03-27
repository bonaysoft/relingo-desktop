package dal

import (
	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
	"github.com/bonaysoft/relingo-desktop/pkg/dal/query"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(dsn string) (*query.Query, error) {
	gormdb, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, err
	}

	if err := gormdb.AutoMigrate(model.Word{}); err != nil {
		return nil, err
	}

	return query.Use(gormdb.Debug()), nil

}
