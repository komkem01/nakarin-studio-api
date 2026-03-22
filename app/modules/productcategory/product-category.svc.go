package productcategory

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
	entitiesinf "nakarin-studio/app/modules/entities/inf"
	"nakarin-studio/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	db     entitiesinf.ProductCategoryEntity
}

type Config struct{}

type Options struct {
	*config.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.ProductCategoryEntity
}

func newService(opt *Options) *Service {
	return &Service{tracer: opt.tracer, db: opt.db}
}

func (s *Service) Create(ctx context.Context, name string, description *string, isActive bool) (*ent.ProductCategoryEntity, error) {
	return s.db.CreateProductCategory(ctx, name, description, isActive)
}

func (s *Service) List(ctx context.Context, isActive *bool) ([]*ent.ProductCategoryEntity, error) {
	return s.db.ListProductCategories(ctx, isActive)
}

func (s *Service) GetByID(ctx context.Context, id string) (*ent.ProductCategoryEntity, error) {
	return s.db.GetProductCategoryByID(ctx, id)
}

func (s *Service) UpdateByID(ctx context.Context, id string, name *string, description *string, isActive *bool) (*ent.ProductCategoryEntity, error) {
	return s.db.UpdateProductCategoryByID(ctx, id, name, description, isActive)
}

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.db.DeleteProductCategoryByID(ctx, id)
}
