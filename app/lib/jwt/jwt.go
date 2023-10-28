package jwt

import (
	"github.com/pascaldekloe/jwt"
	"log"
	"time"
)

func newClaims() jwt.Claims {
	var claims jwt.Claims
	claims.Subject = "alice"
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	return claims
}

func NewJwt(SetMap map[string]interface{}) (token string, err error) {
	claims := newClaims()
	claims.Set = SetMap
	// issue a JWT
	b, err := claims.EdDSASign([]byte("JWTPrivateKey"))
	return string(b), err
}

func VerifyJWT(token string) (ValueMap map[string]interface{}, err error) {

	claims, err := jwt.EdDSACheck([]byte(token), []byte("JWTPrivateKey"))
	if err != nil {
		log.Print("credentials rejected: ", err)
		return nil, err
	}

	err = claims.AcceptTemporal(time.Now(), time.Second)
	if err != nil {
		log.Print("credential constraints violated: ", err)
		return nil, err
	}

	return claims.Set, nil

}
