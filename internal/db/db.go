package db
import (
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

var DB *pg.DB

func InitDB(database string, user string, password string, url string) {
    log.WithFields(log.Fields{"db": database, "user": user, "url": url}).Info("connecting to database")
	DB = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     url,
	})
    Health()

	log.WithFields(log.Fields{"db": database}).Info("connected to database")
}

func Health() bool{
    _, err := DB.Exec("SELECT 1")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database down")
        return false
	}
    return true
}
