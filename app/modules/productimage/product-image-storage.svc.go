package productimage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type storageClient struct {
	client  *minio.Client
	bucket  string
	baseURL string
	enabled bool
}

const productImageMaxSize int64 = 10 * 1024 * 1024

func newStorageClientFromEnv() *storageClient {
	endpointRaw := strings.TrimSpace(os.Getenv("RAILWAY_STORAGE__ENDPOINT"))
	bucket := strings.TrimSpace(os.Getenv("RAILWAY_STORAGE__BUCKET"))
	accessKey := strings.TrimSpace(os.Getenv("RAILWAY_STORAGE__ACCESS_KEY"))
	secretKey := strings.TrimSpace(os.Getenv("RAILWAY_STORAGE__SECRET_KEY"))

	if endpointRaw == "" || bucket == "" || accessKey == "" || secretKey == "" {
		return &storageClient{enabled: false}
	}

	endpoint, secure, baseURL, err := parseStorageEndpoint(endpointRaw)
	if err != nil {
		return &storageClient{enabled: false}
	}

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		return &storageClient{enabled: false}
	}

	return &storageClient{client: cli, bucket: bucket, baseURL: baseURL, enabled: true}
}

func parseStorageEndpoint(raw string) (string, bool, string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", false, "", err
	}
	if u.Scheme == "" {
		return raw, false, strings.TrimRight("https://"+raw, "/"), nil
	}
	endpoint := u.Host
	secure := strings.EqualFold(u.Scheme, "https")
	base := strings.TrimRight(u.Scheme+"://"+u.Host, "/")
	return endpoint, secure, base, nil
}

func (s *storageClient) uploadProductImage(ctx context.Context, productID string, fileHeader *multipart.FileHeader) (string, error) {
	if s == nil || !s.enabled {
		return "", fmt.Errorf("storage is not configured")
	}
	if fileHeader == nil {
		return "", fmt.Errorf("image file is required")
	}
	if fileHeader.Size <= 0 {
		return "", fmt.Errorf("image file is empty")
	}
	if fileHeader.Size > productImageMaxSize {
		return "", fmt.Errorf("image file exceeds 10MB limit")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return "", fmt.Errorf("invalid file type")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(io.LimitReader(file, productImageMaxSize+1))
	if err != nil {
		return "", err
	}
	if int64(len(data)) > productImageMaxSize {
		return "", fmt.Errorf("image file exceeds 10MB limit")
	}

	contentType := http.DetectContentType(data)
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/webp" {
		return "", fmt.Errorf("invalid file content type")
	}

	ext = path.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("product-images/%s/%d-%s%s", productID, time.Now().Unix(), uuid.NewString(), ext)

	_, err = s.client.PutObject(ctx, s.bucket, objectKey, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", s.baseURL, s.bucket, objectKey), nil
}

func (s *storageClient) displayImageURL(ctx context.Context, rawURL string) string {
	if s == nil || !s.enabled {
		return rawURL
	}

	objectKey, ok := s.extractObjectKey(rawURL)
	if !ok {
		return rawURL
	}

	presigned, err := s.client.PresignedGetObject(ctx, s.bucket, objectKey, 24*time.Hour, nil)
	if err != nil {
		return rawURL
	}

	return presigned.String()
}

func (s *storageClient) extractObjectKey(rawURL string) (string, bool) {
	trimmed := strings.TrimSpace(rawURL)
	if trimmed == "" {
		return "", false
	}

	if !strings.Contains(trimmed, "://") {
		return strings.TrimPrefix(trimmed, "/"), true
	}

	u, err := url.Parse(trimmed)
	if err != nil {
		return "", false
	}

	pathValue := strings.TrimPrefix(u.Path, "/")
	if strings.HasPrefix(pathValue, s.bucket+"/") {
		return strings.TrimPrefix(pathValue, s.bucket+"/"), true
	}

	return "", false
}
