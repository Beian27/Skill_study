package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JSON-WEB-TOKEN
// Header Claims Signature

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("zhuzhuzhu")

	c := MyClaims{
		Username: "zhy",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 7,
			Issuer: "zhy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		return
	}
	fmt.Println(signedString)
	time.Sleep(6 * time.Second)
	fmt.Println("---------------------------")
	claims, err1 := jwt.ParseWithClaims(signedString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(claims)
	username := claims.Claims.(*MyClaims).Username
	fmt.Println(username)

}
