package booking

import (
	"nakarin-studio/app/modules/entities"
	entitiesinf "nakarin-studio/app/modules/entities/inf"
	"nakarin-studio/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer      trace.Tracer
	db          entitiesinf.BookingEntity
	detailDB    entitiesinf.BookingDetailEntity
	itemDB      entitiesinf.BookingItemEntity
	paymentDB   entitiesinf.PaymentEntity
	statusLogDB entitiesinf.BookingStatusLogEntity
	txDB        *entities.Service
}

type Config struct{}

type Options struct {
	*config.Config[Config]
	tracer      trace.Tracer
	db          entitiesinf.BookingEntity
	detailDB    entitiesinf.BookingDetailEntity
	itemDB      entitiesinf.BookingItemEntity
	paymentDB   entitiesinf.PaymentEntity
	statusLogDB entitiesinf.BookingStatusLogEntity
	txDB        *entities.Service
}

func newService(opt *Options) *Service {
	return &Service{
		tracer:      opt.tracer,
		db:          opt.db,
		detailDB:    opt.detailDB,
		itemDB:      opt.itemDB,
		paymentDB:   opt.paymentDB,
		statusLogDB: opt.statusLogDB,
		txDB:        opt.txDB,
	}
}
