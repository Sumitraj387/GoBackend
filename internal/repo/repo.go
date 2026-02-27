package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	Db     *gorm.DB
	Logger *logrus.Entry
}
