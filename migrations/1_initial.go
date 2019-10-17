package migrations

import (
	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	pb "github.com/senexi/go-proto-micro/generated/partners"
	log "github.com/sirupsen/logrus"
)

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		for _, model := range []interface{}{(*pb.Partner)(nil)} {
			database := db.(*pg.DB)
			err := database.CreateTable(model, &orm.CreateTableOptions{
				IfNotExists: true,
			})
			if err != nil {
				log.WithFields(log.Fields{"db": db}).Error(err)
			}
		}
		return nil
	}, func(db migrations.DB) error {
		for _, model := range []interface{}{(*pb.Partner)(nil)} {
			database := db.(*pg.DB)
			err := database.DropTable(model, &orm.DropTableOptions{})
			if err != nil {
				log.WithFields(log.Fields{"db": db}).Error(err)
			}
		}
		return nil
	})
}

func Ready() {
	log.Info("initialize migrations")
}
