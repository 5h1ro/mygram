package helpers

import "github.com/dgrijalva/jwt-go"

var secret = "aezakmi"

func Generate(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := parseToken.SignedString([]byte(secret))

	return token
}
