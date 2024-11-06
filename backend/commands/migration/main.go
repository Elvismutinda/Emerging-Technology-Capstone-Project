package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"backend/models"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func RunMigrations() error {
	modelsMigrated := []any{
		models.User{},
		models.Category{},
		models.Budget{},
		models.Transaction{},
		models.Report{},
	}

	queries, err := gormschema.New("postgres").Load(modelsMigrated...)
	if err != nil {
		logrus.Error("Error loading modelsMigrated: ", err)
		return err
	}

	_, err = io.WriteString(os.Stdout, queries)
	if err != nil {
		return err
	}

	logrus.Info("migration complete")
	return nil
}

func main() {
	err := RunMigrations()
	if err != nil {
		logrus.WithField("Error", err).Fatal(err)
	}
}
