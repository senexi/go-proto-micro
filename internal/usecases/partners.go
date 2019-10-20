package usecases

import (
	"github.com/senexi/go-proto-micro/generated/partners"
)

type db interface {
	GetPartners() *partners.PartnerList
}

type PartnerUsecase struct {
	DB db
}

func (p *PartnerUsecase) GetPartners() *partners.PartnerList {
	return p.DB.GetPartners()
}
