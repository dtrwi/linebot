package handler

import (
	"github.com/dtrwi/linebot/internal/app/service"
	"github.com/dtrwi/linebot/internal/pkg"
)

type Option struct {
	pkg.Option
	*service.Service
}
