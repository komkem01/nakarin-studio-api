package payment

import (
	entitiesinf "nakarin-studio/app/modules/entities/inf"
	"nakarin-studio/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer  trace.Tracer
	db      entitiesinf.PaymentEntity
	booking entitiesinf.BookingEntity
	detail  entitiesinf.BookingDetailEntity
	storage *storageClient
}

type Config struct{}

type Options struct {
	*config.Config[Config]
	tracer  trace.Tracer
	db      entitiesinf.PaymentEntity
	booking entitiesinf.BookingEntity
	detail  entitiesinf.BookingDetailEntity
	storage *storageClient
}

func newService(opt *Options) *Service {
	return &Service{tracer: opt.tracer, db: opt.db, booking: opt.booking, detail: opt.detail, storage: opt.storage}
}
