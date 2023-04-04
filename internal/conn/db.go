package conn

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/internal/config"
)

// DB holds the database instance
var db *gorm.DB

// Ping tests if db connection is alive
func Ping() error {
	return db.Exec("SELECT 'DBD::Pg ping test';").Error
}

// Connect sets the db client of database using configuration cfg
func Connect(cfg *config.Database) error {
	host := cfg.Host
	if cfg.Port != 0 {
		host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	}
	uri := url.URL{
		Scheme: "postgres",
		Host:   host,
		Path:   cfg.Name,
		User:   url.UserPassword(cfg.Username, cfg.Password),
	}
	// open a database connection using gorm ORM
	d, err := gorm.Open(postgres.Open(uri.String()), &gorm.Config{})
	if err != nil {
		return err
	}
	db = d

	//AutoMigrate

	makeMigration := Migrate()

	if makeMigration {
		if err := db.AutoMigrate(&domain.User{}); err != nil {
			log.Fatalln(err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	log.Println(sqlDB)
	return nil
}

// DefaultDB returns default db
func DefaultDB() *gorm.DB {
	return db.Debug()
}

// ConnectDB sets the db client of database using default configuration file
func ConnectDB() error {
	cfg := config.DB()
	connectionRenew() //start a connection re-newer
	return Connect(cfg)
}

func connectionRenew() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for t := range ticker.C {
			if err := Ping(); err != nil {
				log.Printf("error: %v [re-connecting database]", err.Error())
				err := Connect(config.DB())
				if err != nil {
					return
				}
				_ = t
			}
		}
	}()
}
