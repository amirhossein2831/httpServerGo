package Auth

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte(config.GetInstance().Get("JWT_SECRET_KEY"))

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 48).Unix(),
			"nbf":   time.Now().Unix(),
			"iat":   time.Now().Unix(),
			"iss":   config.GetInstance().Get("JWT_ISSUER"),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func RetainClaim(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}
