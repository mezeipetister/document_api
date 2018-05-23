package controllers

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func createToken(UID, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UID":  UID,
		"role": role,
	})
	tokenString, errJWT := token.SignedString([]byte("HelloBello"))
	if errJWT != nil {
		fmt.Println(errJWT)
		return ""
	}
	return tokenString
}

// ValidateToken ...
func ValidateToken(tokenString string) (bool, map[string]interface{}) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("HelloBello"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims
	}
	return false, map[string]interface{}{}
}
