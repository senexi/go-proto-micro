package service

import (
	"context"
	"github.com/senexi/go-proto-micro/generated/api"
	"github.com/senexi/go-proto-micro/internal/providers/db"
	"github.com/senexi/go-proto-micro/internal/usecases"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Service implmenting the partner service
type PartnerService struct {
	partnerUsecase *usecases.PartnerUsecase
}

func NewPartnerService() *PartnerService {
	partnerUsecase := &usecases.PartnerUsecase{
		DB: db.NewDatabase(),
	}

	partnerService := &PartnerService{
		partnerUsecase: partnerUsecase,
	}

	return partnerService
}

func (s *PartnerService) GetPartners(ctx context.Context, in *proto.SearchRequest) (*proto.PartnerList, error) {
	partnerList, err := s.partnerUsecase.GetPartners()
	return partnerList, err
}

func (s *PartnerService) AddPartner(ctx context.Context, req *proto.Partner) (*proto.AddPartnerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (s *PartnerService) SearchPartner(ctx context.Context, req *proto.SearchRequest) (*proto.PartnerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPartner not implemented")
}
