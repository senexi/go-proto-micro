package usecases

import (
	"reflect"
	"testing"

	"github.com/senexi/go-proto-micro-generated/generated/api/language/go"
)

type MockDB struct {
}

func (db *MockDB) GetPartners() (*proto.PartnerList, error){
    var partnerList *proto.PartnerList
    partnerList = new(proto.PartnerList)

	partner := proto.Partner{
		Name: "Paul",
	}

    partnerList.Partners = append(partnerList.Partners, &partner)
    return partnerList, nil
}

func TestPartnerUsecase_GetPartners(t *testing.T) {
	type fields struct {
		DB db
	}
	tests := []struct {
		name    string
		fields  fields
		want    *proto.PartnerList
		wantErr bool
	}{
        name: "GetPartners_1",
        fields: {
            DB: MockDB{},
        },
        want: &proto.PartnerList{
            Partners: []proto.Partner{
                proto.
            }
        }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PartnerUsecase{
				DB: tt.fields.DB,
			}
			got, err := p.GetPartners()
			if (err != nil) != tt.wantErr {
				t.Errorf("PartnerUsecase.GetPartners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartnerUsecase.GetPartners() = %v, want %v", got, tt.want)
			}
		})
	}
}
