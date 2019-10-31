package usecases

import (
	"github.com/senexi/go-proto-micro/generated/api"
)
type PartnerUsecase struct {
}

func (p *PartnerUsecase) GetPartners() (*proto.PartnerList, error) {
    return &proto.PartnerList{}, nil
}
