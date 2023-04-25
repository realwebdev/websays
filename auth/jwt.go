package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/realwebdev/clockify/conf"
)

func GenerateJWT(useremail string) (string, error) {
	Accesskey := conf.Conf{}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = useremail
	claims["expirey"] = time.Now().Add(time.Minute * 15).Unix()

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tok.SignedString([]byte(Accesskey.ACCESS))
	if err != nil {
		log.Printf("Somethign went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func RefreshJWT(email string) (string, error) {
	RefreshKey := conf.Conf{}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["expirey"] = time.Now().Add(time.Minute * 200).Unix()

	tokenString, err := token.SignedString([]byte(RefreshKey.ACCESS))
	if err != nil {
		log.Printf("Somethign went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func AuthenticateToken(r *http.Request) error {
	token, err := VerifyToken(r)

	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	conf := conf.Conf{}
	tokenString, err := ExtractToken(r)

	if err != nil {

		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unauthorize signin: %v", token.Header["alg"])
		}

		return []byte(conf.ACCESS), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	StrArr := strings.Split(bearToken, " ") //slice substring

	if len(StrArr) != 2 {

		return "", errors.New("Unable to fetch token from the request")
	}

	return StrArr[1], nil
}
