package memberaddress

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, memberID string, firstName string, lastName *string, phone string, no *string, village *string, street *string, provinceID string, districtID string, subDistrictID string, zipcodeID string) (*ent.MemberAddressEntity, error) {
	return s.db.CreateMemberAddress(ctx, memberID, firstName, lastName, phone, no, village, street, provinceID, districtID, subDistrictID, zipcodeID)
}
