package bookingstatuslog

import (
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

func New(conf *config.Config[Config], db entitiesinf.BookingStatusLogEntity) *Module {
	tracer := otel.Tracer("nakarin-studio.modules.bookingstatuslog")
	svc := newService(&Options{Config: conf, tracer: tracer, db: db})
	return &Module{tracer: tracer, Svc: svc, Ctl: newController(tracer, svc)}
}
