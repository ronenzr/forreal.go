package httpclient

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ronenzr/forreal.go/util/HttpUtil"
	"github.com/ronenzr/forreal.go/util/JwtUtil"
)

const forRealUrl = "http://206.189.96.100/api/generate"

func GetCommand(phrase string) string {

	values := map[string]string{"phrase": phrase, "tone": "rewrite"}
	res := HttpUtil.PostWithBody(forRealUrl, values)

	if res == nil {
		fmt.Println("Failed to fetch data from server, please contact ForReal administrator")
		return ""
	}

	resultToken := res["result"]
	verifyOptions := jwt.MapClaims{
		"iss": "ForReal.API",
		"sub": phrase,
		"aud": "ForReal-Client"}

	claims := JwtUtil.ParseAndValidate(resultToken.(string), verifyOptions)

	if claims == nil {
		fmt.Println("Response authentication failed, please contact ForReal administrator")
		return ""
	}

	return claims["text"].(string)
}
