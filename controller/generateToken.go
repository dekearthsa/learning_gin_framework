package controller

import (
	"learning_gin_framework/model"
	"time"

	"github.com/golang-jwt/jwt"
)

const SIGNATURE string = "kd#isji#@0nivok@--sdf8#83845@@##d2#+#=+#dsjif"

func GenterateToken(Email string) string {
	expirationTime := time.Now().Add(5 * time.Minute)

	warpData := &model.StructEncodeToken{
		Email: Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, warpData)
	enCodeToken, err := token.SignedString([]byte(SIGNATURE))
	if err != nil {
		return err.Error()
	}
	return enCodeToken
}
