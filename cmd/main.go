package main

import (
	"context"
	"github.com/Thunderbirrd/exchange-backend/internal/config"
	"github.com/Thunderbirrd/exchange-backend/internal/repository"
	"github.com/Thunderbirrd/exchange-backend/internal/repository/postgres"
	"github.com/Thunderbirrd/exchange-backend/internal/service"
	handler "github.com/Thunderbirrd/exchange-backend/internal/transport/http"
	"github.com/Thunderbirrd/exchange-backend/pkg"
	"github.com/Thunderbirrd/exchange-backend/pkg/utils"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Exchange app
// @version 1.0
// @description Exchange app

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	conf := config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	utils.EnvToString(&conf.Password, "DB_PASSWORD", "qwerty")

	db, err := postgres.NewPostgresDB(conf)

	if err != nil {
		logrus.Fatalf("failed to inizialed db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(pkg.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
