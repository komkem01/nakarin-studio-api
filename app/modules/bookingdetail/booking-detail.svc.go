package bookingdetail

import (
	entitiesinf "nakarin-studio/app/modules/entities/inf"
	"nakarin-studio/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	db     entitiesinf.BookingDetailEntity
}

type Config struct{}

type Options struct {
	*config.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.BookingDetailEntity
}

func newService(opt *Options) *Service { return &Service{tracer: opt.tracer, db: opt.db} }
