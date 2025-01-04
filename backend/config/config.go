package config

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type LocalConfig struct {
	JWTSecretKey string `required:"true" split_words:"true"`
	PostgresDSN  string `required:"true" split_words:"true"`
	Port         string `required:"true" split_words:"true"`
	DBHost       string `required:"false" split_words:"true"`
	Environment  string `required:"false" split_words:"true"`
	DockerDSN    string `required:"false" split_words:"true"`
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

func ConnectDB() (db *gorm.DB, err error) {
	// get env variables
	err = godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file", err)
	}

	localCfg, err := FromEnv()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	dsn := localCfg.PostgresDSN

	if localCfg.Environment == "docker" {
		dsn = localCfg.DockerDSN // use docker's dsn where hostname is the db container name
	}

	// Connect to the database
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		logrus.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	// Ping the database to ensure connectivity
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Error initializing GORM with *sql.DB: %v", err)
		return nil, err
	}

	// Assign db to global DB variable
	DB = db
	return db, nil
}
