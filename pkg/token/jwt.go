package token

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	issuer string
	secret string
}

func NewJwtService() *jwtService {
	secret := os.Getenv("JWT_SECRET_KEY")
	issuer := "upl"

	return &jwtService{issuer: issuer, secret: secret}
}

type JwtClaim struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

func (j *jwtService) GenerateToken(accountID, role string) (string, error) {
	tokenDuration := time.Duration(time.Minute * 60 * 24 * 7) //7 days

	claims := &JwtClaim{
		accountID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(tokenDuration).Unix(),
			Issuer:    "upl",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *jwtService) isTokenValid(t *jwt.Token) (interface{}, error) {
	if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, fmt.Errorf("invalid token %v", t)
	}

	return []byte(j.secret), nil
}

func (j *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, j.isTokenValid)

	return err == nil
}

type payload struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
}

func (j *jwtService) RetriveTokenPayload(token string) (*payload, error) {
	t, err := jwt.Parse(token, j.isTokenValid)
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("unable to parse jwt claims")
	}

	payload := &payload{
		AccountID: claims["account_id"].(string),
		Role:      claims["role"].(string),
	}

	return payload, err
}
