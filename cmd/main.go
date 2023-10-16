package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/api"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/config"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/logger"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage"
)

func main() {
	logger.Init()
	log := logger.GetLogger()
	log.Info("logger initialized")

	cfg := config.Load()
	log.Info("config initialized")
	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		fmt.Println(psqlUrl)
		log.Fatalf("failed to connect to database: %v, %s", err, psqlUrl)
	}

	strg := storage.NewStoragePg(psqlConn)

	apiServer := api.New(api.RoutetOptions{
		Cfg:     &cfg,
		Storage: strg,
		Log:     log,
	})


	if err = apiServer.Run(fmt.Sprintf(":%s", cfg.HttpPort)); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
