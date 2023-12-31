package postgres

import (
	"fmt"
	"tgotify/config"
	models "tgotify/storage"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Gormdb is a database wrapper that uses GORM for interacting with a PostgreSQL database.
type Gormdb struct {
	*gorm.DB
}

// New creates a new Gormdb instance and initializes the database connection using the provided configuration.
func New(config *config.StorageConfig) *Gormdb {
	db := &Gormdb{connect(config)}
	db.sync()
	return db
}

// connect establishes a new database connection based on the provided configuration and returns the GORM DB instance.
func connect(config *config.StorageConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
		"disable",
	)

	// Open a connection to the PostgreSQL database using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("unable to connect to the database")
	}
	logrus.Info("successfully connected to the database")
	return db
}

// sync ensures that the necessary database tables are created or updated to match the defined GORM models.
func (db *Gormdb) sync() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Chat{})
}
