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
