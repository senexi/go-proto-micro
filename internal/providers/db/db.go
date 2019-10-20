package db

import (
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

var postgresDB *pg.DB
var isConnected bool

type DB struct {

}

func Connect(database string, user string, password string, url string) {
    log.WithFields(log.Fields{"db": database, "user": user, "url": url}).Info("connecting to database")
	postgresDB = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     url,
    })
	Health()

	log.WithFields(log.Fields{"db": database}).Info("connected to database")
}

func NewDatabase() *DB{
    return &DB{}
}

func Health() bool {
	_, err := postgresDB.Exec("SELECT 1")
	if err != nil {
        log.WithFields(log.Fields{"error": err}).Error("database down")
        isConnected = false
		return false
    }
    isConnected = true
	return isConnected
}
