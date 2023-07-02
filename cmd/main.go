package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kazakhattila/kapster"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	/*
		// dal'we -> coffee shop rob it.
		// as well as REST API youtube video -> ROB + hardcore
		// as well as ROB 3 medium messenger apps + make your version in the end...
		//
				1) main get and rob it to the fullest
				2) services + handlers rob to the fullest 
				3) rob the riches from the repos in the end rob everything nahooi! + write git mistake + gopath
				4) rob docker do talova
				5) rob everything, lead to the docker
				6) then to solve problems: refresh, my database specifics, my getters, Channels and my problems...
				7) then proceed to rob leetcode today -> if can't solve or too lazy to write and decide -> rob others...
				8)
				9)
	*/
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error loading database: %s", err.Error())

	}

	repos := repositories.NewRepository(db)
	service := services.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(kapster.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print(" Shutting down an app ")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
