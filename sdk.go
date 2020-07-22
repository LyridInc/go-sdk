package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LyridInc/go-sdk/client"
	"github.com/LyridInc/go-sdk/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/tatsushid/go-fastping"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path"
	"sync"
	"time"
)

type LyridClient struct {
	mux sync.Mutex

	isUploading bool
	token       string
	lyridurl    string
	lyridaccess string
	lyridsecret string
}

var instance *LyridClient
var once sync.Once

func GetInstance() *LyridClient {
	once.Do(func() {
		instance = &LyridClient{}
	})
	return instance
}

func (client *LyridClient) GetLyridURL() string {
	if client.lyridurl == "" {
		client.lyridurl = "api.lyrid.io"
	}

	return client.lyridurl
}

func (client *LyridClient) SetLyridURL(url string) {
	client.lyridurl = url
}

func pinglyridurl() error {

	client := GetInstance()
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", client.lyridurl)
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

func (client *LyridClient) Initialize(access string, secret string) error {
	// ping the
	err := pinglyridurl()

	if err != nil {
		fmt.Println(err)
		return err
	}

	// use
	client.lyridaccess = access
	client.lyridsecret = secret

	_, err = client.login()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(client.token)

	//validatetoken()
	return nil
}

func (lc *LyridClient) GetUserProfile() *model.LyridUser {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/user")
		if err == nil {
			if response.StatusCode == 200 {
				var user model.LyridUser
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &user)
				return &user
			}
		}
	}

	return nil
}

func (lc *LyridClient) GetAccountProfile() []*model.Account {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/accounts")
		if err == nil {
			if response.StatusCode == 200 {
				var accounts []*model.Account
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &accounts)
				return accounts
			}
		}
	}

	return nil
}

func (lc *LyridClient) GetApps() []*model.App {
	// tbd
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/serverless/app/get")
		if err == nil {
			if response.StatusCode == 200 {
				var apps []*model.App
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &apps)
				return apps
			}
		}
	}
	return nil
}

func (lc *LyridClient) GetModules(AppId string) []*model.Module {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/serverless/app/get/" + AppId)
		if err == nil {
			if response.StatusCode == 200 {
				var modules []*model.Module
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &modules)
				return modules
			}
		}
	}
	return nil
}

func (lc *LyridClient) GetRevisions(AppId string, ModuleId string) []*model.ModuleRevision {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/serverless/app/get/" + AppId + "/" + ModuleId)
		if err == nil {
			if response.StatusCode == 200 {
				var revisions []*model.ModuleRevision
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &revisions)
				return revisions
			}
		}
	}
	return nil
}

func (lc *LyridClient) GetFunctions(AppId string, ModuleId string, RevisionId string) []*model.Function {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("api/serverless/app/get/" + AppId + "/" + ModuleId + "/" + RevisionId)
		if err == nil {
			if response.StatusCode == 200 {
				var functions []*model.Function
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &functions)
				return functions
			}
		}
	}
	return nil
}

func (lc *LyridClient) ExecuteFunction(FunctionId string, Framework string, Body string) interface{} {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Post("api/serverless/app/execute/"+FunctionId+"/"+Framework, Body)
		if err == nil {
			if response.StatusCode == 200 {
				var retValue interface{}
				databyte, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(databyte, &retValue)
				return retValue
			}
		}
	}

	return nil
}

func (lc *LyridClient) GetAccountPolicies() []*model.Policy {
	return nil
}

func (lc *LyridClient) GetModulePolicies(ModuleId string) []*model.Policy {
	return nil
}

func (client *LyridClient) checktoken() bool {
	if client.istokenexpired() {
		client.login()
		return client.istokenexpired()
	}

	return false
}

func (client *LyridClient) istokenexpired() bool {
	token, _, err := new(jwt.Parser).ParseUnverified(client.token, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

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

func (client *LyridClient) login() (string, error) {
	jsonData := map[string]string{"key": client.lyridaccess, "secret": client.lyridsecret}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://"+path.Join(client.GetLyridURL(), "auth"), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println("The HTTP request to Lyrid Server failed.\n", err)
		return "", err
	} else {
		databyte, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			var config model.UserAccessToken
			json.Unmarshal(databyte, &config)
			client.token = config.Token
			return config.Token, nil
		} else {
			log.Println("The HTTP request to Lyrid Server failed.\n", string(databyte))
			return "", errors.New(string(databyte))
		}
	}
}
