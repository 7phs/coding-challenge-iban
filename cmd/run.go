package cmd

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/db"
	"github.com/7phs/coding-challenge-iban/logger"
	"github.com/7phs/coding-challenge-iban/model"
	"github.com/7phs/coding-challenge-iban/restapi"
	"github.com/7phs/coding-challenge-iban/restapi/handler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.ParseConfig()

		logger.Init(conf)

		log.Info(ApplicationInfo())

		if err := conf.Validate(); err!=nil {
			log.Fatal("config: invalid, ", err)
		}

		if err := db.Init(conf); err!=nil {
			log.Fatal("db: failed to init, ", err)
		}

		restapi.Init(conf)

		model.Init(&model.Dependencies{
			CountriesFormat: db.CountriesFmt,
		})

		server := restapi.
			NewServer(conf, handler.DefaultRouter(conf)).
			Run()

		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt)
		<-stop
		log.Info("interrupt signal")

		server.Shutdown()

		db.Shutdown()
		log.Info("finish")
	},
}
