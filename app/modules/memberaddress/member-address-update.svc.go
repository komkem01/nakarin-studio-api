package memberaddress

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, memberID *string, firstName *string, lastName *string, phone *string, no *string, village *string, street *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string) (*ent.MemberAddressEntity, error) {
	return s.db.UpdateMemberAddressByID(ctx, id, memberID, firstName, lastName, phone, no, village, street, provinceID, districtID, subDistrictID, zipcodeID)
}
