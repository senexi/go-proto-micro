package service

import (
	"context"
	"github.com/senexi/go-proto-micro-generated/generated/api/language/go"
	"github.com/senexi/go-proto-micro/internal/providers/db"
	"github.com/senexi/go-proto-micro/internal/usecases"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Service implmenting the partner service
type MockPartnerService struct {
	partnerUsecase *usecases.PartnerUsecase
}

func NewMockPartnerService() *PartnerService {
	partnerUsecase := &usecases.PartnerUsecase{
		DB: db.NewMockDatabase(),
	}

	partnerService := &PartnerService{
		partnerUsecase: partnerUsecase,
	}

	return partnerService
}

func (s *MockPartnerService) GetPartners(ctx context.Context, in *proto.SearchRequest) (*proto.PartnerList, error) {
	partnerList, err := s.partnerUsecase.GetPartners()
	return partnerList, err
}

func (s *MockPartnerService) AddPartner(ctx context.Context, req *proto.Partner) (*proto.AddPartnerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (s *MockPartnerService) SearchPartner(ctx context.Context, req *proto.SearchRequest) (*proto.PartnerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPartner not implemented")
}
