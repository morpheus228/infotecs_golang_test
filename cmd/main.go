package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	wallet "github.com/morpheus228/infotecs_golang_test"
	"github.com/morpheus228/infotecs_golang_test/pkg/handler"
	"github.com/morpheus228/infotecs_golang_test/pkg/repository"
	"github.com/morpheus228/infotecs_golang_test/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := intiConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(wallet.Server)

	go func() { srv.Run(viper.GetString("port"), handlers.InitRoutes()) }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	srv.Shutdown(context.Background())
	db.Close()
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
