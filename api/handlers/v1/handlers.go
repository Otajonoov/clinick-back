package handlers

import (
	"gitlab.com/QuvonchbekOtajonov/clinic-back/config"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/logger"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	Storage storage.StorageI
	log     logger.Logger
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage *storage.StorageI
	Log     logger.Logger
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		Storage: *options.Storage,
		log:     options.Log,
	}
}
