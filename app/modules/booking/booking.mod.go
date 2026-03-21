package booking

import (
	"nakarin-studio/app/modules/entities"
	entitiesinf "nakarin-studio/app/modules/entities/inf"
	"nakarin-studio/internal/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Module struct {
	tracer trace.Tracer
	Svc    *Service
	Ctl    *Controller
}

func New(conf *config.Config[Config], db entitiesinf.BookingEntity, detailDB entitiesinf.BookingDetailEntity, itemDB entitiesinf.BookingItemEntity, paymentDB entitiesinf.PaymentEntity, statusLogDB entitiesinf.BookingStatusLogEntity, txDB *entities.Service) *Module {
	tracer := otel.Tracer("nakarin-studio.modules.booking")
	svc := newService(&Options{Config: conf, tracer: tracer, db: db, detailDB: detailDB, itemDB: itemDB, paymentDB: paymentDB, statusLogDB: statusLogDB, txDB: txDB})
	return &Module{tracer: tracer, Svc: svc, Ctl: newController(tracer, svc)}
}
