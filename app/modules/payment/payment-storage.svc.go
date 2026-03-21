package payment

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

const proofMaxSize int64 = 5 * 1024 * 1024

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
		// MinIO client expects endpoint without scheme.
		return raw, false, strings.TrimRight("https://"+raw, "/"), nil
	}
	endpoint := u.Host
	secure := strings.EqualFold(u.Scheme, "https")
	base := strings.TrimRight(u.Scheme+"://"+u.Host, "/")
	return endpoint, secure, base, nil
}

func (s *storageClient) uploadProof(ctx context.Context, paymentID string, fileHeader *multipart.FileHeader) (string, error) {
	if s == nil || !s.enabled {
		return "", fmt.Errorf("storage is not configured")
	}
	if fileHeader == nil {
		return "", fmt.Errorf("proof file is required")
	}
	if fileHeader.Size <= 0 {
		return "", fmt.Errorf("proof file is empty")
	}
	if fileHeader.Size > proofMaxSize {
		return "", fmt.Errorf("proof file exceeds 5MB limit")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".pdf" {
		return "", fmt.Errorf("invalid file type")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(io.LimitReader(file, proofMaxSize+1))
	if err != nil {
		return "", err
	}
	if int64(len(data)) > proofMaxSize {
		return "", fmt.Errorf("proof file exceeds 5MB limit")
	}

	contentType := http.DetectContentType(data)
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "application/pdf" {
		return "", fmt.Errorf("invalid file content type")
	}

	ext = path.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("payment-proofs/%s/%d-%s%s", paymentID, time.Now().Unix(), uuid.NewString(), ext)

	_, err = s.client.PutObject(ctx, s.bucket, objectKey, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", s.baseURL, s.bucket, objectKey), nil
}
