package payment

import (
	"context"
	"fmt"
)

func (s *Service) UploadProof(ctx context.Context, id string, proofURL string) error {
	proof := proofURL
	if proof == "" {
		return fmt.Errorf("proof_url is required")
	}
	_, err := s.db.UpdatePaymentByID(ctx, id, nil, nil, nil, nil, nil, &proof, nil, nil)
	return err
}

func (s *Service) Approve(ctx context.Context, id string, note *string) error {
	status := "paid"
	_, err := s.db.UpdatePaymentByID(ctx, id, nil, nil, nil, nil, &status, nil, note, nil)
	return err
}

func (s *Service) Reject(ctx context.Context, id string, reason *string) error {
	status := "failed"
	_, err := s.db.UpdatePaymentByID(ctx, id, nil, nil, nil, nil, &status, nil, reason, nil)
	return err
}
