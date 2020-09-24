package main

import (
	"log"

	"github.com/alandtsang/easyjwt"
)

func main() {
	data := map[string]interface{}{
		"name":  "Your name",
		"age":   18,
		"admin": true,
	}

	token, err := easyjwt.GenerateToken(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token)

	got, err := easyjwt.ParseToken(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", got)
}
