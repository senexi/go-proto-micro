package service

import (
    "context"
    "github.com/senexi/go-proto-micro/internal/usecases"
    "github.com/senexi/go-proto-micro/internal/providers/db"
	pb "github.com/senexi/go-proto-micro/generated/partners"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Service implmenting the partner service
type PartnerService struct {
    partnerUsecase *usecases.PartnerUsecase
}

func NewPartnerService() *PartnerService{
    partnerUsecase := &usecases.PartnerUsecase{
        DB: db.NewDatabase(),
    }

    partnerService := &PartnerService{
        partnerUsecase: partnerUsecase,
    }

    return partnerService
}

func (s *PartnerService) GetPartners(ctx context.Context, in *pb.SearchRequest) (*pb.PartnerList, error) {
	var partnerList *pb.PartnerList
	partnerList = s.partnerUsecase.GetPartners()
	return partnerList, nil
}

func (s *PartnerService) AddPartner(ctx context.Context, req *pb.Partner) (*pb.AddPartnerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (s *PartnerService) SearchPartner(ctx context.Context, req *pb.SearchRequest) (*pb.PartnerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPartner not implemented")
}
