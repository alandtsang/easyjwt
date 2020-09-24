package jwtdemo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alandtsang/easyjwt"
)

var data = map[string]interface{}{
	"name":  "Alan",
	"id":    1,
	"admin": true,
}

func TestGenerateToken(t *testing.T) {
	token, err := easyjwt.GenerateToken(data)
	assert.Equal(t, nil, err)
	t.Log(token)
}

func TestGenerateCustomToken(t *testing.T) {
	var (
		secret = "This is a simple secret"
		expire = time.Minute * 30
	)

	token, err := easyjwt.GenerateCustomToken(data, secret, expire)
	assert.Equal(t, nil, err)
	t.Log(token)
}

func TestParseNormalToken(t *testing.T) {
	token, err := easyjwt.GenerateToken(data)
	assert.Equal(t, nil, err)

	got, err := easyjwt.ParseToken(token)
	assert.Equal(t, nil, err)

	gotData := got.(map[string]interface{})
	for k, v := range data {
		if k == "id" {
			assert.Equal(t, v, int(gotData[k].(float64)))
			continue
		}
		assert.Equal(t, v, gotData[k])
	}
}

func TestParseExpiredToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjAwOTM2NzM1LCJuYW1lIjoiQWxhbiJ9.ZNDmS_5qHDS_DhgU6ICFup4Z72pgnLl7QZv5mFTtm8o"
	_, err := easyjwt.ParseToken(token)
	assert.Equal(t, easyjwt.ErrExpiredToken, err)
}

func TestParseSignatureInvalidToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.pazba9Pj009HgANP4pTCQAHpXNU7pVbjIGff_plktSzsa9rXTGzFngaawzXGEO6Q0Hx5dtGi-dMDlIadV81o3Q"
	_, err := easyjwt.ParseToken(token)
	assert.Equal(t, easyjwt.ErrSignatureInvalid, err)
}

func TestParseCustomToken(t *testing.T) {
	var (
		secret = "This is a simple secret"
		expire = time.Minute * 30
	)

	token, err := easyjwt.GenerateCustomToken(data, secret, expire)
	assert.Equal(t, nil, err)

	got, err := easyjwt.ParseCustomToken(token, secret)
	assert.Equal(t, nil, err)

	gotData := got.(map[string]interface{})
	for k, v := range data {
		if k == "id" {
			assert.Equal(t, v, int(gotData[k].(float64)))
			continue
		}
		assert.Equal(t, v, gotData[k])
	}
}
