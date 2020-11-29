package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Id  uint   `json:"id"`
}
type JWT struct {
	Header    Header
	Payload   Payload
	Signature string
	Token     string
}

func NewJWT(id uint) JWT {
	var jwt JWT
	jwt.Header = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jwt.Payload = Payload{
		Iss: "redrock",
		Exp: strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat: strconv.FormatInt(time.Now().Unix(), 10),
		Id:  id,
	}
	h, _ := json.Marshal(jwt.Header)
	p, _ := json.Marshal(jwt.Payload)
	baseh := base64.StdEncoding.EncodeToString(h)
	basep := base64.StdEncoding.EncodeToString(p)
	secret := baseh + "." + basep
	key := "redrock"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(secret))
	s := mac.Sum(nil)
	jwt.Signature = base64.StdEncoding.EncodeToString(s)
	jwt.Token = secret + "." + jwt.Signature
	return jwt
}

func Check(token string) (jwt JWT, err error) {
	err = errors.New("token error")
	arr := strings.Split(token, ".")
	if len(arr) < 3 {
		fmt.Println("59------", err)
		return
	}
	baseh := arr[0]
	h, err := base64.StdEncoding.DecodeString(baseh)
	if err != nil {
		fmt.Println("68-------", err)
		return
	}
	err = json.Unmarshal(h, &jwt.Header)
	if err != nil {
		fmt.Println("73-------", err)
		return
	}
	basep := arr[1]
	p, err := base64.StdEncoding.DecodeString(basep)
	if err != nil {
		fmt.Println("79-------", err)
		return
	}
	err = json.Unmarshal(p, &jwt.Payload)
	if err != nil {
		fmt.Println("84-------", err)
		return
	}
	bases := arr[2]
	s1, err := base64.StdEncoding.DecodeString(bases)
	if err != nil {
		fmt.Println("84-------", err)
		return
	}
	se := baseh + "." + basep
	w := []byte("redrock")
	mac := hmac.New(sha256.New, w)
	mac.Write([]byte(se))
	s2 := mac.Sum(nil)
	if string(s1) != string(s2) {
		return
	} else {
		jwt.Signature = arr[2]
		jwt.Token = token
	}
	return jwt, nil
}
