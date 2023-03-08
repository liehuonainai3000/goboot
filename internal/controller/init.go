package controller

import (
	"github.com/liehuonainai3000/goboot/frame/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger = zap.L().Sugar()

var binder *utils.GinBinder = utils.NewGinBinder()
