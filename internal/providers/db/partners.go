package db

import(
    "github.com/senexi/go-proto-micro/generated/api"
    log "github.com/sirupsen/logrus"
)

func (db *DB) GetPartners() (*partners.PartnerList, error){
    var result []*partners.Partner
    err := db.db.Model(result).Limit(1000).Select()
    if err != nil {
        log.Error(err)
    }

    partnerList := &partners.PartnerList{
        Partners: result,
    }
    return partnerList, err
}