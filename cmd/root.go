package cmd

import (
	"github.com/dtrwi/linebot/config"
	"github.com/dtrwi/linebot/internal/app/repository"
	"github.com/dtrwi/linebot/internal/app/server"
	"github.com/dtrwi/linebot/internal/app/service"
	"github.com/dtrwi/linebot/internal/pkg"
)

// Execute adds all child commands.
// This is called by main.main(). It only needs to happen once.
func Execute() {
	// var err error
	cfg := config.Config()
	// driver := driver.NewDriver(cfg)

	// var postgreSQL *sqlx.DB
	// if cfg.GetBool("postgres.is_enabled") {
	// 	postgreSQL, err = driver.GetPostgreSQLConn()
	// 	if err != nil {
	// 		logrus.Fatalf("failed to start, error connect to PostgreSQL | +v", err)
	// 		return
	// 	}
	// 	defer postgreSQL.Close()
	// }

	option := pkg.Option{
		Config: cfg,
		// PostgreSQL:    postgreSQL,
	}

	repository := wiringRepository(repository.Option{
		Option: option,
	})

	service := wiringService(service.Option{
		Option:     option,
		Repository: repository,
	})

	server := server.NewServer(option, service)
	server.StartApp()
}

func wiringRepository(option repository.Option) *repository.Repository {
	// wiring up all repositories here

	repo := repository.Repository{}
	return &repo
}

func wiringService(option service.Option) *service.Service {
	// wiring up all services

	svc := service.Service{}
	return &svc
}
