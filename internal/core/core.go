package core

import (
	"gobackend/internal/repo"

	"github.com/sirupsen/logrus"
)

type ICore interface {
}
type Core struct {
	Logger *logrus.Entry
	RepoV1 repo.Repository
}
