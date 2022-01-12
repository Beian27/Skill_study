package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JSON-WEB-TOKEN
// Header Claims Signature

func main() {
	mySigningKey := []byte("zhuzhuzhu")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Unix() + 5,
		"iss":      "zhy",
		"nbf":      time.Now().Unix() - 5,
		"username": "zhy",
	})
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		return
	}
	fmt.Println(signedString)
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------------")
	claims, err1 := jwt.ParseWithClaims(signedString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(claims)
	username := claims.Claims.(*jwt.MapClaims)
	fmt.Println(username)
}
