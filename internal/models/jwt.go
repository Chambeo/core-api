package models

import (
	customErrors "chambeo-core-api/internal/errors"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	Email  string
	Role   string
	UserID string
	jwt.RegisteredClaims
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(email string, role string, userID string) (signedToken string, err error) {
	claims := &JwtClaim{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours))},
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

//ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New(customErrors.UnmarshallClaimsError)
		return
	}

	if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now().Local()) {
		err = errors.New(customErrors.ExpiredJWTError)
		return
	}

	return

}
