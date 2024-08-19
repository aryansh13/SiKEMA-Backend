package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Uid    string
	Type   string
	TypeID string
}

func GenerateToken(claim JWTClaim) (string, error) {
	token_lifetime, _ := strconv.Atoi("3600")

	claims := jwt.MapClaims{}
	claims["uid"] = claim.Uid
	claims["type"] = claim.Type
	claims["type_id"] = claim.TypeID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(token_lifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("cfiomaciojcihomwciohb3dcwriou3r"))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("cfiomaciojcihomwciohb3dcwriou3r"), nil
	})
}

func ValidateToken(tokenString string) error {
	token, err := ParseToken(tokenString)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return fmt.Errorf("Unauthenticated")
}

func ExtractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ExtractTokenClaim(tokenString string) (JWTClaim, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return JWTClaim{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var claim JWTClaim
		// fmt.Println("Before converted : " + claims["uid"].(string))
		uid, _ := claims["uid"].(string)
		utype, _ := claims["type"].(string)
		typeID, _ := claims["type_id"].(string)
		claim.Uid = uid
		claim.Type = utype
		claim.TypeID = typeID
		// fmt.Println("After converted : " + claim.Uid)
		return claim, nil
	}
	return JWTClaim{}, nil
}
