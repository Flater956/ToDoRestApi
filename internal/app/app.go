package app

import (
	"ToDoRestApi/internal/repository"
	"ToDoRestApi/internal/server"
	"ToDoRestApi/internal/service"
	"ToDoRestApi/internal/transfer/http"
	"ToDoRestApi/pkg/database/postgresdb"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env %s", err.Error())
	}

	db, err := initDB()

	if err != nil {
		log.Fatalf("error initializing DB %s", err.Error())
	}

	rep := repository.NewRepository(db)
	serv := service.NewService(rep)
	handlers := http.NewHandler(serv)

	server := new(server.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	return postgresdb.NewPostgresDB(postgresdb.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
		SSLMode: viper.GetString("db.sslmode"),
	})
}
