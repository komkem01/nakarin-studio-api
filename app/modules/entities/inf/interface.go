package entitiesinf

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

// ObjectEntity defines the interface for object entity operations such as create, retrieve, update, and soft delete.
type ExampleEntity interface {
	CreateExample(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
	GetExampleByID(ctx context.Context, id uuid.UUID) (*ent.Example, error)
	UpdateExampleByID(ctx context.Context, id uuid.UUID, status ent.ExampleStatus) (*ent.Example, error)
	SoftDeleteExampleByID(ctx context.Context, id uuid.UUID) error
	ListExamplesByStatus(ctx context.Context, status ent.ExampleStatus) ([]*ent.Example, error)
}
type ExampleTwoEntity interface {
	CreateExampleTwo(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
}

type GenderEntity interface {
	CreateGender(ctx context.Context, name string, isActive bool) (*ent.GenderEntity, error)
	GetGenderByID(ctx context.Context, id string) (*ent.GenderEntity, error)
	UpdateGenderByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.GenderEntity, error)
	ListGenders(ctx context.Context, isActive *bool) ([]*ent.GenderEntity, error)
	DeleteGenderByID(ctx context.Context, id string) error
}

type PrefixEntity interface {
	CreatePrefix(ctx context.Context, genderID string, name string, isActive bool) (*ent.PrefixEntity, error)
	GetPrefixByID(ctx context.Context, id string) (*ent.PrefixEntity, error)
	UpdatePrefixByID(ctx context.Context, id string, genderID *string, name *string, isActive *bool) (*ent.PrefixEntity, error)
	ListPrefixes(ctx context.Context, genderID *string, isActive *bool) ([]*ent.PrefixEntity, error)
	DeletePrefixByID(ctx context.Context, id string) error
}

type ProvinceEntity interface {
	CreateProvince(ctx context.Context, name string, isActive bool) (*ent.ProvinceEntity, error)
	GetProvinceByID(ctx context.Context, id string) (*ent.ProvinceEntity, error)
	UpdateProvinceByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.ProvinceEntity, error)
	ListProvinces(ctx context.Context, isActive *bool) ([]*ent.ProvinceEntity, error)
	DeleteProvinceByID(ctx context.Context, id string) error
}

type DistrictEntity interface {
	CreateDistrict(ctx context.Context, provinceID string, name string, isActive bool) (*ent.DistrictEntity, error)
	GetDistrictByID(ctx context.Context, id string) (*ent.DistrictEntity, error)
	UpdateDistrictByID(ctx context.Context, id string, provinceID *string, name *string, isActive *bool) (*ent.DistrictEntity, error)
	ListDistricts(ctx context.Context, provinceID *string, isActive *bool) ([]*ent.DistrictEntity, error)
	DeleteDistrictByID(ctx context.Context, id string) error
}

type SubDistrictEntity interface {
	CreateSubDistrict(ctx context.Context, districtID string, name string, isActive bool) (*ent.SubDistrictEntity, error)
	GetSubDistrictByID(ctx context.Context, id string) (*ent.SubDistrictEntity, error)
	UpdateSubDistrictByID(ctx context.Context, id string, districtID *string, name *string, isActive *bool) (*ent.SubDistrictEntity, error)
	ListSubDistricts(ctx context.Context, districtID *string, isActive *bool) ([]*ent.SubDistrictEntity, error)
	DeleteSubDistrictByID(ctx context.Context, id string) error
}

type ZipcodeEntity interface {
	CreateZipcode(ctx context.Context, subDistrictID string, name string, isActive bool) (*ent.ZipcodeEntity, error)
	GetZipcodeByID(ctx context.Context, id string) (*ent.ZipcodeEntity, error)
	UpdateZipcodeByID(ctx context.Context, id string, subDistrictID *string, name *string, isActive *bool) (*ent.ZipcodeEntity, error)
	ListZipcodes(ctx context.Context, subDistrictID *string, isActive *bool) ([]*ent.ZipcodeEntity, error)
	DeleteZipcodeByID(ctx context.Context, id string) error
}

type BookingEntity interface {
	CreateBooking(ctx context.Context, bookingNo string, status *string, payment *string) (*ent.BookingEntity, error)
	GetBookingByID(ctx context.Context, id string) (*ent.BookingEntity, error)
	UpdateBookingByID(ctx context.Context, id string, bookingNo *string, status *string, payment *string) (*ent.BookingEntity, error)
	ListBookings(ctx context.Context, status *string, payment *string) ([]*ent.BookingEntity, error)
	DeleteBookingByID(ctx context.Context, id string) error
}

type BookingDetailEntity interface {
	CreateBookingDetail(ctx context.Context, bookingID string, firstName string, lastName *string, phone string) (*ent.BookingDetailEntity, error)
	GetBookingDetailByID(ctx context.Context, id string) (*ent.BookingDetailEntity, error)
	UpdateBookingDetailByID(ctx context.Context, id string, bookingID *string, firstName *string, lastName *string, phone *string) (*ent.BookingDetailEntity, error)
	ListBookingDetails(ctx context.Context, bookingID *string) ([]*ent.BookingDetailEntity, error)
	DeleteBookingDetailByID(ctx context.Context, id string) error
}

type MemberEntity interface {
	CreateMember(ctx context.Context, genderID string, prefixID *string, role *string, firstName string, lastName *string, phone string) (*ent.MemberEntity, error)
	GetMemberByID(ctx context.Context, id string) (*ent.MemberEntity, error)
	UpdateMemberByID(ctx context.Context, id string, genderID *string, prefixID *string, role *string, firstName *string, lastName *string, phone *string) (*ent.MemberEntity, error)
	ListMembers(ctx context.Context, genderID *string, prefixID *string, role *string, phone *string) ([]*ent.MemberEntity, error)
	DeleteMemberByID(ctx context.Context, id string) error
}

type MemberBookingEntity interface {
	CreateMemberBooking(ctx context.Context, memberID string, bookingID string) (*ent.MemberBookingEntity, error)
	GetMemberBookingByID(ctx context.Context, id string) (*ent.MemberBookingEntity, error)
	UpdateMemberBookingByID(ctx context.Context, id string, memberID *string, bookingID *string) (*ent.MemberBookingEntity, error)
	ListMemberBookings(ctx context.Context, memberID *string, bookingID *string) ([]*ent.MemberBookingEntity, error)
	DeleteMemberBookingByID(ctx context.Context, id string) error
}

type MemberAddressEntity interface {
	CreateMemberAddress(ctx context.Context, memberID string, firstName string, lastName *string, phone string, no *string, village *string, street *string, provinceID string, districtID string, subDistrictID string, zipcodeID string) (*ent.MemberAddressEntity, error)
	GetMemberAddressByID(ctx context.Context, id string) (*ent.MemberAddressEntity, error)
	UpdateMemberAddressByID(ctx context.Context, id string, memberID *string, firstName *string, lastName *string, phone *string, no *string, village *string, street *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string) (*ent.MemberAddressEntity, error)
	ListMemberAddresses(ctx context.Context, memberID *string, provinceID *string, districtID *string, subDistrictID *string, zipcodeID *string, phone *string) ([]*ent.MemberAddressEntity, error)
	DeleteMemberAddressByID(ctx context.Context, id string) error
}
