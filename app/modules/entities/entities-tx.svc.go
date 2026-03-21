package entities

import (
	"context"

	"github.com/uptrace/bun"
)

func (s *Service) RunInTx(ctx context.Context, fn func(ctx context.Context, txSvc *Service) error) error {
	return s.root.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		txSvc := &Service{db: tx, root: s.root}
		return fn(ctx, txSvc)
	})
}
