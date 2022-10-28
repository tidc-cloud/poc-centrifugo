package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func generateJWT() (string, error) {

	//https://blog.logrocket.com/jwt-authentication-go/#generating-jwts-using-golang-jwt-package
	var sampleSecretKey = []byte("my_secret")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(1000000 * time.Hour).Unix()
	claims["authorized"] = true
	claims["sub"] = "1234"
	claims["channels"] = [2]string{"activity:a_trail03"} //"#1234",
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	jwt, _ := generateJWT()
	fmt.Println("jwt", jwt)

	// var message Message
	// if err != nil {
	// 	return
	// }
	// err = json.NewEncoder(writer).Encode(message)
	// if err != nil {
	// 	return
	// }

}
