package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/LyridInc/go-sdk/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/tatsushid/go-fastping"
	"io/ioutil"
	"net"
	"net/http"
	"path"
	"time"
)

var (
	lyridurl    = "app.lyrid.io"
	lyridaccess = ""
	lyridsecret = ""

	lyridtoken = ""
)

func Hello() string {
	return "Hello, world."
}

func GetLyridURL() string {
	return lyridurl
}

func SetLyridURL(url string) {
	lyridurl = url
}

func pinglyridurl() error {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", lyridurl)
	if err != nil {
		return err
	}
	p.AddIPAddr(ra)
	err = p.Run()
	if err != nil {
		return err
	}

	return nil
}

func Initialize(access string, secret string) error {
	// ping the
	//err := pinglyridurl()

	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}

	// use
	lyridaccess = access
	lyridsecret = secret

	_, err := login()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(lyridtoken)

	//validatetoken()
	return nil
}

func GetUserProfile() *model.LyridUser {
	return nil
}

func checktoken() bool {
	if istokenexpired() {
		login()
		return istokenexpired()
	}

	return false
}

func istokenexpired() bool {
	token, _, err := new(jwt.Parser).ParseUnverified(lyridtoken, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["exp"])
		var tm time.Time
		switch iat := claims["iat"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0)
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)
		}

		if time.Now().After(tm) {
			return true
		}
	} else {
		fmt.Println(err)
	}

	return false
}

func login() (string, error) {

	jsonData := map[string]string{"key": lyridaccess, "secret": lyridsecret}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://"+path.Join(GetLyridURL(), "auth"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request to Lyrid Server failed. %s\n", err)
		return "", err
	} else {
		var config model.UserAccessToken
		databyte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(databyte, &config)

		lyridtoken = config.Token

		return config.Token, nil
	}
}
