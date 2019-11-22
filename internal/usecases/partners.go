package usecases

import (
	"github.com/senexi/go-proto-micro-generated/generated/api/language/go"
)

type db interface {
	GetPartners() (*proto.PartnerList, error)
}

type PartnerUsecase struct {
	DB db
}

func (p *PartnerUsecase) GetPartners() (*proto.PartnerList, error) {
	return p.DB.GetPartners()
}
