package main

import (
	"fmt"
	"gophercises/secrets_api"
	"log"
)

func main() {
	v := secrets_api.NewVault("super_secret", "secrets")
	err := v.Add("aws", "aws_key")
	if err != nil {
		log.Fatal(err)
	}
	val, err := v.Get("aws")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

}
