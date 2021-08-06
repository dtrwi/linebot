package service

import (
	"github.com/dtrwi/linebot/internal/app/repository"
	"github.com/dtrwi/linebot/internal/pkg"
)

type Option struct {
	pkg.Option
	*repository.Repository
}

type Service struct {
}
