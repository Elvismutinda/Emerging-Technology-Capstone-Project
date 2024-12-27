package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type LocalConfig struct {
	JWTSecretKey string `required:"true" split_words:"true"`
	PostgresDSN  string `required:"true" split_words:"true"`
	Port         string `required:"true" split_words:"true"`
}

func FromEnv() (cfg *LocalConfig, err error) {
	fromFileToEnv()

	cfg = &LocalConfig{}

	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Debug("No config files found to load to env. Defaulting to environment.")
	} else {
		logrus.Debug("Config files found loading to env.")
	}
}

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file", err)
	}

	localCfg, err := FromEnv()
	if err != nil {
		logrus.Error(err)
		return
	}

	dsn := localCfg.PostgresDSN

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Error connecting to database", err)
	}
	DB = db
	logrus.Info("Connected to database")
}
