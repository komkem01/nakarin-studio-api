package admin

import (
	"fmt"
	"os"
	"time"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	AdminID  string `json:"admin_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	TokenUse string `json:"token_use"`
	jwt.RegisteredClaims
}

func (s *Service) generateAccessToken(admin *ent.AdminEntity) (string, error) {
	return s.generateToken(admin, "access", 2*time.Hour)
}

func (s *Service) generateRefreshToken(admin *ent.AdminEntity) (string, error) {
	return s.generateToken(admin, "refresh", 30*24*time.Hour)
}

func (s *Service) generateToken(admin *ent.AdminEntity, tokenUse string, ttl time.Duration) (string, error) {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY is required")
	}

	now := time.Now()
	claims := &Claims{
		AdminID:  admin.ID.String(),
		Username: admin.Username,
		Role:     "admin",
		TokenUse: tokenUse,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   admin.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func (s *Service) ValidateAccessToken(raw string) (*Claims, error) {
	return s.validateToken(raw, "access")
}

func (s *Service) ValidateRefreshToken(raw string) (*Claims, error) {
	return s.validateToken(raw, "refresh")
}

func (s *Service) validateToken(raw string, expectedTokenUse string) (*Claims, error) {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY is required")
	}

	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims.Role != "admin" {
		return nil, fmt.Errorf("invalid role")
	}
	if claims.TokenUse != expectedTokenUse {
		return nil, fmt.Errorf("invalid token use")
	}

	return claims, nil
}
