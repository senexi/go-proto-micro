package db

import(
    "github.com/senexi/go-proto-micro/generated/partners"
)

func getPartners() *partners.PartnerList{
    var partnerList *partners.PartnerList
    partnerList = new(partners.PartnerList)

	partner := partners.Partner{
		Name: "Paul",
	}

    partnerList.Partners = append(partnerList.Partners, &partner)
    return partnerList
}