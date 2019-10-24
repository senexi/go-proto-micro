package db

import(
    "github.com/senexi/go-proto-micro/generated/api"
    log "github.com/sirupsen/logrus"
)

func (db *DB) GetPartners() (*proto.PartnerList, error){
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