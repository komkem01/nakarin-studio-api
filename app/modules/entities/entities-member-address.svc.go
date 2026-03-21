package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateMemberAddress(ctx context.Context, memberID string, firstName string, lastName *string, phone string, no *string, village *string, street *string, provinceID string, districtID string, subDistrictID string, zipcodeID string) (*ent.MemberAddressEntity, error) {
	memberUUID, provinceUUID, districtUUID, subDistrictUUID, zipcodeUUID, err := parseAddressRequiredUUIDs(memberID, provinceID, districtID, subDistrictID, zipcodeID)
	if err != nil {
		return nil, err
	}

	firstNameValue := strings.TrimSpace(firstName)
	if firstNameValue == "" {
		return nil, fmt.Errorf("first_name is required")
	}

	phoneValue := strings.TrimSpace(phone)
	if phoneValue == "" {
		return nil, fmt.Errorf("phone is required")
	}

	model := &ent.MemberAddressEntity{
		ID:            uuid.New(),
		MemberID:      memberUUID,
		FirstName:     firstNameValue,
		LastName:      normalizeOptionalString(lastName),
		Phone:         phoneValue,
		No:            normalizeOptionalString(no),
		Village:       normalizeOptionalString(village),
		Street:        normalizeOptionalString(street),
		ProvinceID:    provinceUUID,
		DistrictID:    districtUUID,
		SubDistrictID: subDistrictUUID,
		ZipcodeID:     zipcodeUUID,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetMemberAddressByID(ctx context.Context, id string) (*ent.MemberAddressEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberAddressEntity{}
	if err := s.db.NewSelect().Model(model).Where("member_address.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListMemberAddresses(ctx context.Context, memberID *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string, phone *string) ([]*ent.MemberAddressEntity, error) {
	items := make([]*ent.MemberAddressEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("member_address.created_at DESC")

	if memberID != nil && strings.TrimSpace(*memberID) != "" {
		memberUUID, err := uuid.Parse(strings.TrimSpace(*memberID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_address.member_id = ?", memberUUID)
	}
	if provinceID != nil && strings.TrimSpace(*provinceID) != "" {
		provinceUUID, err := uuid.Parse(strings.TrimSpace(*provinceID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_address.province_id = ?", provinceUUID)
	}
	if districtID != nil && strings.TrimSpace(*districtID) != "" {
		districtUUID, err := uuid.Parse(strings.TrimSpace(*districtID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_address.district_id = ?", districtUUID)
	}
	if subDistrictID != nil && strings.TrimSpace(*subDistrictID) != "" {
		subDistrictUUID, err := uuid.Parse(strings.TrimSpace(*subDistrictID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_address.sub_district_id = ?", subDistrictUUID)
	}
	if zipcodeID != nil && strings.TrimSpace(*zipcodeID) != "" {
		zipcodeUUID, err := uuid.Parse(strings.TrimSpace(*zipcodeID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_address.zipcode_id = ?", zipcodeUUID)
	}
	if phone != nil && strings.TrimSpace(*phone) != "" {
		q = q.Where("member_address.phone ILIKE ?", "%"+strings.TrimSpace(*phone)+"%")
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateMemberAddressByID(ctx context.Context, id string, memberID *string, firstName *string, lastName *string, phone *string, no *string, village *string, street *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string) (*ent.MemberAddressEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberAddressEntity{}
	if err := s.db.NewSelect().Model(model).Where("member_address.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if memberID != nil {
		memberUUID, err := uuid.Parse(strings.TrimSpace(*memberID))
		if err != nil {
			return nil, err
		}
		model.MemberID = memberUUID
	}
	if firstName != nil {
		firstNameValue := strings.TrimSpace(*firstName)
		if firstNameValue == "" {
			return nil, fmt.Errorf("first_name is required")
		}
		model.FirstName = firstNameValue
	}
	if lastName != nil {
		model.LastName = normalizeOptionalString(lastName)
	}
	if phone != nil {
		phoneValue := strings.TrimSpace(*phone)
		if phoneValue == "" {
			return nil, fmt.Errorf("phone is required")
		}
		model.Phone = phoneValue
	}
	if no != nil {
		model.No = normalizeOptionalString(no)
	}
	if village != nil {
		model.Village = normalizeOptionalString(village)
	}
	if street != nil {
		model.Street = normalizeOptionalString(street)
	}
	if provinceID != nil {
		provinceUUID, err := uuid.Parse(strings.TrimSpace(*provinceID))
		if err != nil {
			return nil, err
		}
		model.ProvinceID = provinceUUID
	}
	if districtID != nil {
		districtUUID, err := uuid.Parse(strings.TrimSpace(*districtID))
		if err != nil {
			return nil, err
		}
		model.DistrictID = districtUUID
	}
	if subDistrictID != nil {
		subDistrictUUID, err := uuid.Parse(strings.TrimSpace(*subDistrictID))
		if err != nil {
			return nil, err
		}
		model.SubDistrictID = subDistrictUUID
	}
	if zipcodeID != nil {
		zipcodeUUID, err := uuid.Parse(strings.TrimSpace(*zipcodeID))
		if err != nil {
			return nil, err
		}
		model.ZipcodeID = zipcodeUUID
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("member_id", "first_name", "last_name", "phone", "no", "village", "street", "province_id", "district_id", "sub_district_id", "zipcode_id", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteMemberAddressByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.MemberAddressEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func parseAddressRequiredUUIDs(memberID string, provinceID string, districtID string, subDistrictID string, zipcodeID string) (uuid.UUID, uuid.UUID, uuid.UUID, uuid.UUID, uuid.UUID, error) {
	memberUUID, err := uuid.Parse(strings.TrimSpace(memberID))
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, err
	}
	provinceUUID, err := uuid.Parse(strings.TrimSpace(provinceID))
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, err
	}
	districtUUID, err := uuid.Parse(strings.TrimSpace(districtID))
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, err
	}
	subDistrictUUID, err := uuid.Parse(strings.TrimSpace(subDistrictID))
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, err
	}
	zipcodeUUID, err := uuid.Parse(strings.TrimSpace(zipcodeID))
	if err != nil {
		return uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, uuid.Nil, err
	}
	return memberUUID, provinceUUID, districtUUID, subDistrictUUID, zipcodeUUID, nil
}
