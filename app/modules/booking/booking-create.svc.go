package booking

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingNo string, status *string, payment *string, cancelledReason *string, internalNote *string, trackingAttemptCount *int, lastTrackingAt *time.Time, deliveryMemberAddressID *string, deliveryFirstName *string, deliveryLastName *string, deliveryPhone *string, deliveryNo *string, deliveryVillage *string, deliveryStreet *string, deliveryProvinceID *string, deliveryDistrictID *string, deliverySubDistrictID *string, deliveryZipcodeID *string, deliveryNote *string) (*ent.BookingEntity, error) {
	return s.db.CreateBooking(ctx, bookingNo, status, payment, cancelledReason, internalNote, trackingAttemptCount, lastTrackingAt, deliveryMemberAddressID, deliveryFirstName, deliveryLastName, deliveryPhone, deliveryNo, deliveryVillage, deliveryStreet, deliveryProvinceID, deliveryDistrictID, deliverySubDistrictID, deliveryZipcodeID, deliveryNote)
}
