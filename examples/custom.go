package main

import (
	"github.com/alandtsang/easyjwt"
	"log"
	"time"
)

func main() {
	var (
		secret = "custom"
		expire = time.Minute * 30
		data   = map[string]interface{}{
			"name":  "Your name",
			"age":   18,
			"admin": true,
		}
	)
	token, err := easyjwt.GenerateCustomToken(data, secret, expire)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token)

	got, err := easyjwt.ParseCustomToken(token, secret)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", got)
}
