package memberaddress

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, memberID *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string, phone *string) ([]*ent.MemberAddressEntity, error) {
	return s.db.ListMemberAddresses(ctx, memberID, provinceID, districtID, subDistrictID, zipcodeID, phone)
}
