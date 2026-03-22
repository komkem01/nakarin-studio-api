package payment

import (
	"context"
	"fmt"
)

func (s *Service) ProofViewURL(ctx context.Context, id string) (string, error) {
	item, err := s.db.GetPaymentByID(ctx, id)
	if err != nil {
		return "", err
	}
	if item == nil || item.ProofURL == nil || *item.ProofURL == "" {
		return "", fmt.Errorf("proof url not found")
	}

	if s.storage == nil {
		return *item.ProofURL, nil
	}

	return s.storage.proofViewURL(ctx, *item.ProofURL)
}
