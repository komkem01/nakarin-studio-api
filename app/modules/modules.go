package modules

import (
	"log/slog"
	"sync"

	"nakarin-studio/app/modules/booking"
	"nakarin-studio/app/modules/bookingdetail"
	"nakarin-studio/app/modules/district"
	"nakarin-studio/app/modules/entities"
	"nakarin-studio/app/modules/example"
	"nakarin-studio/app/modules/gender"
	"nakarin-studio/app/modules/prefix"
	"nakarin-studio/app/modules/province"
	"nakarin-studio/app/modules/sentry"
	"nakarin-studio/app/modules/specs"
	"nakarin-studio/app/modules/subdistrict"
	"nakarin-studio/app/modules/zipcode"
	"nakarin-studio/internal/config"
	"nakarin-studio/internal/database"
	"nakarin-studio/internal/log"
	"nakarin-studio/internal/otel/collector"

	exampletwo "nakarin-studio/app/modules/example-two"

	appConf "nakarin-studio/config"
	// "nakarin-studio/app/modules/kafka"
)

type Modules struct {
	Conf          *config.Module[appConf.Config]
	Specs         *specs.Module
	Log           *log.Module
	OTEL          *collector.Module
	Sentry        *sentry.Module
	DB            *database.DatabaseModule
	ENT           *entities.Module
	Gender        *gender.Module
	Prefix        *prefix.Module
	Province      *province.Module
	District      *district.Module
	SubDistrict   *subdistrict.Module
	Zipcode       *zipcode.Module
	Booking       *booking.Module
	BookingDetail *bookingdetail.Module
	// Kafka *kafka.Module
	Example  *example.Module
	Example2 *exampletwo.Module
}

func modulesInit() {
	confMod := config.New(&appConf.App)
	specsMod := specs.New(config.Conf[specs.Config](confMod.Svc))
	conf := confMod.Svc.Config()

	logMod := log.New(config.Conf[log.Option](confMod.Svc))
	otel := collector.New(config.Conf[collector.Config](confMod.Svc))
	log := log.With(slog.String("module", "modules"))

	sentryMod := sentry.New(config.Conf[sentry.Config](confMod.Svc))

	db := database.New(conf.Database.Sql)
	entitiesMod := entities.New(db.Svc.DB())
	genderMod := gender.New(config.Conf[gender.Config](confMod.Svc), entitiesMod.Svc)
	prefixMod := prefix.New(config.Conf[prefix.Config](confMod.Svc), entitiesMod.Svc)
	provinceMod := province.New(config.Conf[province.Config](confMod.Svc), entitiesMod.Svc)
	districtMod := district.New(config.Conf[district.Config](confMod.Svc), entitiesMod.Svc)
	subDistrictMod := subdistrict.New(config.Conf[subdistrict.Config](confMod.Svc), entitiesMod.Svc)
	zipcodeMod := zipcode.New(config.Conf[zipcode.Config](confMod.Svc), entitiesMod.Svc)
	bookingMod := booking.New(config.Conf[booking.Config](confMod.Svc), entitiesMod.Svc)
	bookingDetailMod := bookingdetail.New(config.Conf[bookingdetail.Config](confMod.Svc), entitiesMod.Svc)
	exampleMod := example.New(config.Conf[example.Config](confMod.Svc), entitiesMod.Svc)
	exampleMod2 := exampletwo.New(config.Conf[exampletwo.Config](confMod.Svc), entitiesMod.Svc)
	// kafka := kafka.New(&conf.Kafka)
	mod = &Modules{
		Conf:          confMod,
		Specs:         specsMod,
		Log:           logMod,
		OTEL:          otel,
		Sentry:        sentryMod,
		DB:            db,
		ENT:           entitiesMod,
		Gender:        genderMod,
		Prefix:        prefixMod,
		Province:      provinceMod,
		District:      districtMod,
		SubDistrict:   subDistrictMod,
		Zipcode:       zipcodeMod,
		Booking:       bookingMod,
		BookingDetail: bookingDetailMod,
		Example:       exampleMod,
		Example2:      exampleMod2,
	}

	log.Infof("all modules initialized")
}

var (
	once sync.Once
	mod  *Modules
)

func Get() *Modules {
	once.Do(modulesInit)

	return mod
}
