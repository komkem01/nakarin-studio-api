package payment

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
)

var ErrProofFileInvalid = errors.New("invalid proof file")

func (s *Service) UploadProof(ctx context.Context, id string, fileHeader *multipart.FileHeader) (string, error) {
	if s.storage == nil {
		return "", fmt.Errorf("payment storage client is not initialized")
	}

	proof, err := s.storage.uploadProof(ctx, id, fileHeader)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrProofFileInvalid, err)
	}

	_, err = s.db.UpdatePaymentByID(ctx, id, nil, nil, nil, nil, nil, &proof, nil, nil)
	if err != nil {
		return "", err
	}

	return proof, nil
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
