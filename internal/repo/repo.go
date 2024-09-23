package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRepository interface {
}
type Repository struct {
	Db     *gorm.DB
	Logger *logrus.Entry
}
