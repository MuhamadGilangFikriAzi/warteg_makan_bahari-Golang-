package authenticator

import (
	"WMB/delivery/apprequest"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token interface {
	CreateToken(dataLogin apprequest.LoginResponse) (string, error)
	VerifAccessToken(tokenString string) (jwt.MapClaims, error)
}

type TokenConfig struct {
	AplicationName      string
	JwtSignatureKey     string
	JwtSignatureMethod  *jwt.SigningMethodHMAC
	AccessTokenDuration time.Duration
}

type token struct {
	config TokenConfig
}

func NewToken(config TokenConfig) Token {
	return &token{
		config,
	}
}

func (t *token) CreateToken(dataLogin apprequest.LoginResponse) (string, error) {
	//now := time.Now().UTC()
	//fmt.Println(t.config.AccessTokenDuration)
	//end := now.Add(t.config.AccessTokenDuration)
	claims := MyClaims{ // Menyiapkan struct dengan isi yg dibutuhkan
		StandardClaims: jwt.StandardClaims{
			Issuer: t.config.AplicationName,
		},
		Username: dataLogin.Name,
		Email:    dataLogin.Email,
	}
	//claims.IssuedAt = now.Unix()
	//claims.ExpiresAt = end.Unix()
	//jwt.SigningMethodES256
	token := jwt.NewWithClaims(t.config.JwtSignatureMethod, claims) // membuat jwt dengan format method dan claim berupa struct
	return token.SignedString([]byte(t.config.JwtSignatureKey))     // mendapatkan token
}

func (t *token) VerifAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signid Method invalid")
		} else if method != t.config.JwtSignatureMethod {
			return nil, fmt.Errorf("signid Method invalid")
		}
		return []byte(t.config.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
