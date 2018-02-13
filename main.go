package main

import (
	"fmt"
  "github.com/tsunny/go-common.git/utils"
  "github.com/dgrijalva/jwt-go"
)

func main() {

  claims := utils.ParseClaims()
  fmt.Printf("Claims struct: %v\n", claims)

  fmt.Printf("Will try casting jwt.Claims to jwt.MapClaims \n")

  claimsMap := claims.(jwt.MapClaims)

  fmt.Printf("jwt.MapClaims struct: %v\n", claimsMap)

}
