package common

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/edufriendchen/applet-platform/common/tls"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"time"
)

const (
	DefaultMaxOpen     = 10
	DefaultMaxIdle     = 10
	DefaultMaxLifetime = 3
)

type Config struct {
	Host        string
	Port        string
	User        string
	Password    string
	Name        string
	MaxOpen     int
	MaxIdle     int
	MaxLifetime int // in minutes
	MaxIdleTime int // in minutes
	CA          []byte
	ServerName  string
	ParseTime   bool
	Location    string
}

func dataSourceName(config Config) string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	val := url.Values{}

	if config.ParseTime {
		val.Add("parseTime", "1")
	}
	if len(config.Location) > 0 {
		val.Add("loc", config.Location)
	}
	if config.CA != nil {
		val.Add("tls", "custom")
	}

	if len(val) == 0 {
		return connection
	}
	return fmt.Sprintf("%s?%s", connection, val.Encode())
}

// DB return new sql db
func DB(config Config) (*sql.DB, error) {
	if config.CA != nil && len(config.ServerName) != 0 {
		if err := mysql.RegisterTLSConfig("custom", tls.WithServerAndCA(config.ServerName, config.CA)); err != nil {
			log.Println(err)
			return nil, err
		}
	} else if config.CA != nil {
		if err := mysql.RegisterTLSConfig("custom", tls.WithCA(config.CA)); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	db, err := sql.Open("mysql", dataSourceName(config))
	if err != nil {
		return nil, err
	}

	if config.MaxOpen > 0 {
		db.SetMaxOpenConns(config.MaxOpen)
	} else {
		db.SetMaxOpenConns(DefaultMaxOpen)
	}

	if config.MaxIdle > 0 {
		db.SetMaxIdleConns(config.MaxIdle)
	} else {
		db.SetMaxIdleConns(DefaultMaxIdle)
	}

	if config.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Minute)
	} else {
		db.SetConnMaxLifetime(time.Duration(DefaultMaxLifetime) * time.Minute)
	}

	if config.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Minute)
	}

	return db, nil
}

func GetDatabase(prefix string) (*sqlx.DB, error) {
	// dbCa := config.GetBinary(`database.ca`)
	dbServerName := viper.GetString(prefix + `.server.name`)
	dbHost := viper.GetString(prefix + `.host`)
	dbPort := viper.GetString(prefix + `.port`)
	dbUser := viper.GetString(prefix + `.user`)
	dbPass := viper.GetString(prefix + `.pass`)
	dbName := viper.GetString(prefix + `.name`)
	maxOpen := viper.GetInt(prefix + `.max.open`)
	maxIdle := viper.GetInt(prefix + `.max.idle`)
	maxLifetime := viper.GetInt(prefix + `.max.lifetime`)
	maxIdletime := viper.GetInt(prefix + `.max.idletime`)

	if dbHost == "" || dbUser == "" {
		return nil, errors.New("empty credential db")
	}

	dbConfig := Config{
		Host:        dbHost,
		Port:        dbPort,
		User:        dbUser,
		Password:    dbPass,
		Name:        dbName,
		MaxOpen:     int(maxOpen),
		MaxIdle:     int(maxIdle),
		MaxLifetime: int(maxLifetime),
		MaxIdleTime: int(maxIdletime),
		// CA:          dbCa,
		ServerName: dbServerName,
		Location:   "Asia/Jakarta",
		ParseTime:  true,
	}

	db, err := DB(dbConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbx := sqlx.NewDb(db, "mysql")
	return dbx, nil
}
