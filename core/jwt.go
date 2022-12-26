package core

import (
	"fmt"
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func JWTSign(res *SheetsAddUserResponse) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject: strconv.Itoa(res.RowId),
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func JWTVerify(token string) (string, error) {
	t, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			// Validate the alg is HMAC
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return jwtSecret, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", fmt.Errorf("cannot parse claims ok: %v, valid: %v", ok, t.Valid)
	}

	return claims["sub"].(string), nil
}
