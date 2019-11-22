package db

import (
	"github.com/senexi/go-proto-micro-generated/generated/api/language/go"
	log "github.com/sirupsen/logrus"
)

type MockDB struct {
}

func NewMockDatabase() *MockDB {
	log.Info("using mock database")
	return &MockDB{}
}

func (db *MockDB) GetPartners() (*proto.PartnerList, error) {
	var partnerList *proto.PartnerList
	partnerList = new(proto.PartnerList)

	partner := proto.Partner{
		Name: "Paul",
	}

	partnerList.Partners = append(partnerList.Partners, &partner)
	return partnerList, nil
}
