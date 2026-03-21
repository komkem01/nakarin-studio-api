package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateMember(ctx context.Context, genderID string, prefixID *string, role *string, firstName string, lastName *string, phone string) (*ent.MemberEntity, error) {
	genderUUID, err := uuid.Parse(strings.TrimSpace(genderID))
	if err != nil {
		return nil, err
	}

	prefixUUID, err := parseOptionalUUID(prefixID)
	if err != nil {
		return nil, err
	}

	roleValue, err := normalizeMemberRole(role)
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

	model := &ent.MemberEntity{
		ID:        uuid.New(),
		GenderID:  genderUUID,
		PrefixID:  prefixUUID,
		Role:      roleValue,
		FirstName: firstNameValue,
		LastName:  normalizeOptionalString(lastName),
		Phone:     phoneValue,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetMemberByID(ctx context.Context, id string) (*ent.MemberEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberEntity{}
	if err := s.db.NewSelect().Model(model).Where("member.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListMembers(ctx context.Context, genderID *string, prefixID *string, role *string, phone *string) ([]*ent.MemberEntity, error) {
	items := make([]*ent.MemberEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("member.created_at DESC")

	if genderID != nil && strings.TrimSpace(*genderID) != "" {
		genderUUID, err := uuid.Parse(strings.TrimSpace(*genderID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member.gender_id = ?", genderUUID)
	}

	if prefixID != nil {
		prefixUUID, err := parseOptionalUUID(prefixID)
		if err != nil {
			return nil, err
		}
		if prefixUUID == nil {
			q = q.Where("member.prefix_id is null")
		} else {
			q = q.Where("member.prefix_id = ?", *prefixUUID)
		}
	}

	if role != nil {
		roleValue, err := normalizeMemberRole(role)
		if err != nil {
			return nil, err
		}
		q = q.Where("member.role = ?", roleValue)
	}

	if phone != nil && strings.TrimSpace(*phone) != "" {
		q = q.Where("member.phone ILIKE ?", "%"+strings.TrimSpace(*phone)+"%")
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateMemberByID(ctx context.Context, id string, genderID *string, prefixID *string, role *string, firstName *string, lastName *string, phone *string) (*ent.MemberEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberEntity{}
	if err := s.db.NewSelect().Model(model).Where("member.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if genderID != nil {
		genderUUID, err := uuid.Parse(strings.TrimSpace(*genderID))
		if err != nil {
			return nil, err
		}
		model.GenderID = genderUUID
	}

	if prefixID != nil {
		prefixUUID, err := parseOptionalUUID(prefixID)
		if err != nil {
			return nil, err
		}
		model.PrefixID = prefixUUID
	}

	if role != nil {
		roleValue, err := normalizeMemberRole(role)
		if err != nil {
			return nil, err
		}
		model.Role = roleValue
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

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("gender_id", "prefix_id", "role", "first_name", "last_name", "phone", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteMemberByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.MemberEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func normalizeMemberRole(value *string) (ent.MemberRole, error) {
	if value == nil {
		return ent.MemberRoleCustomer, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return "", fmt.Errorf("role is invalid")
	}

	normalized := ent.MemberRole(raw)
	switch normalized {
	case ent.MemberRoleCustomer, ent.MemberRoleAdmin:
		return normalized, nil
	default:
		return "", fmt.Errorf("role is invalid")
	}
}

func parseOptionalUUID(value *string) (*uuid.UUID, error) {
	if value == nil {
		return nil, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return nil, nil
	}

	parsed, err := uuid.Parse(raw)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
