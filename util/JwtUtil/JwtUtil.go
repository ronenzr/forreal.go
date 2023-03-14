package JwtUtil

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ronenzr/forreal.go/util/AuthUtil"
)

func ParseAndValidate(token string, verifyOptions jwt.MapClaims) jwt.MapClaims {

	claims := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(AuthUtil.PublicKey))
	})

	// ... error handling
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if !result.Valid {
		println("security token is not valid, please report this to ForReal package administrator")
	}

	// do something with decoded claims
	for key, val := range claims {

		if verifyOptions[key] != nil && val != verifyOptions[key] {
			println("security token is not valid, please report this to ForReal package administrator")
			return nil
		}
	}

	return claims
}
