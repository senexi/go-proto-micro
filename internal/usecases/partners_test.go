package usecases

import (
	"reflect"
	"testing"

	"github.com/senexi/go-proto-micro/generated/api"
)

type MockDB struct {
}

func (db *MockDB) GetPartners() (*partners.PartnerList, error){
    var partnerList *partners.PartnerList
    partnerList = new(partners.PartnerList)

	partner := partners.Partner{
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
		want    *partners.PartnerList
		wantErr bool
	}{
        name: "GetPartners_1",
        fields: {
            DB: MockDB{},
        },
        want: &partners.PartnerList{
            Partners: []partners.Partner{
                partners.
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
