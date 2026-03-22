package productimage

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"nakarin-studio/app/modules/entities/ent"
)

var ErrProductImageFileInvalid = errors.New("invalid product image file")

func (s *Service) UploadAndCreate(ctx context.Context, productID string, fileHeader *multipart.FileHeader, altText *string, sortOrder int, isActive bool) (*ent.ProductImageEntity, error) {
	if s.storage == nil {
		return nil, fmt.Errorf("product image storage client is not initialized")
	}

	imageURL, err := s.storage.uploadProductImage(ctx, productID, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrProductImageFileInvalid, err)
	}

	item, err := s.db.CreateProductImage(ctx, productID, imageURL, altText, sortOrder, isActive)
	if err != nil {
		return nil, err
	}

	item.ImageURL = s.storage.displayImageURL(ctx, item.ImageURL)
	return item, nil
}
