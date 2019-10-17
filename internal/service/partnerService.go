package service

import (
	"context"

	pb "github.com/senexi/go-proto-micro/generated/partners"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Server implmenting the partner service
type PartnerService struct {
}

func (s *PartnerService) GetPartners(ctx context.Context, in *pb.SearchRequest) (*pb.PartnerList, error) {
	var partnerList *pb.PartnerList
	partnerList = new(pb.PartnerList)

	partner := pb.Partner{
		Name: "Paul",
	}

	partnerList.Partners = append(partnerList.Partners, &partner)
	return partnerList, nil
}

func (s *PartnerService) AddPartner(ctx context.Context, req *pb.Partner) (*pb.AddPartnerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (s *PartnerService) SearchPartner(ctx context.Context, req *pb.SearchRequest) (*pb.PartnerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPartner not implemented")
}
