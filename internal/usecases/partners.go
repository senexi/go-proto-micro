package usecases

import (
	"github.com/senexi/go-proto-micro/generated/api"
)

type db interface {
	GetPartners() (*partners.PartnerList, error)
}

type PartnerUsecase struct {
	DB db
}

func (p *PartnerUsecase) GetPartners() (*partners.PartnerList, error) {
	return p.DB.GetPartners()
}
