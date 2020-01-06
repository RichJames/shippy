package main

import (
	"time"

	pb "github.com/RichJames/shippy/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (

	// Define a secure key string used
	// as a salt when hasing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("yN8MdDTmRsnzGu6L")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User 	*pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(data *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claim
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer: "shippy.user.srv",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
