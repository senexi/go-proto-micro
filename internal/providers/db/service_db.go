package db

import (
	"github.com/go-pg/pg"
	"github.com/senexi/go-proto-micro-generated/generated/api/language/go"
	log "github.com/sirupsen/logrus"
)

type DB struct {
	db *pg.DB
}

func NewDatabase() *DB {
	return &DB{
		db: postgresDB,
	}
}

func (db *DB) GetPartners() (*proto.PartnerList, error) {
	var result []*proto.Partner
	err := db.db.Model(result).Limit(1000).Select()
	if err != nil {
		log.Error(err)
	}

	partnerList := &proto.PartnerList{
		Partners: result,
	}
	return partnerList, err
}
