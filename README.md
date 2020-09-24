# easyjwt

Simple jwt in golang

easyjwt is used to generate jwt token and parse jwt token,
use HS256 to sign the payload, support custom secret and expiration time,
default expiration time is 1 hour.

## Usage

normal:

```go
// Your payload
data := map[string]interface{}{
    "name":  "Your name",
    "age":   18,
    "admin": true,
}

token, err := easyjwt.GenerateToken(data)
if err != nil {
    log.Fatal(err)
}

got, err := easyjwt.ParseToken(token)
if err != nil {
    log.Fatal(err)
}
```

custom:

```go
var (
    // Custom secret string
    secret = "custom"

    // Custom token expiration time
    expire = time.Minute * 30

    // Your payload
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

got, err := easyjwt.ParseCustomToken(token, secret)
if err != nil {
    log.Fatal(err)
}
```

## Get Help

The fastest way to get response is to send email to my mail:

- <zengxianglong0@gmail.com>

## LICENCE

Please refer to [LICENSE](https://github.com/alandtsang/easyjwt/blob/master/LICENSE) file.
