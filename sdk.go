package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/LyridInc/go-sdk/client"
	"github.com/LyridInc/go-sdk/model"
	"github.com/golang-jwt/jwt"
	"github.com/tatsushid/go-fastping"
)

type LyridClient struct {
	mux sync.Mutex

	isUploading bool
	token       string
	lyridurl    string
	lyridaccess string
	lyridsecret string

	simulateserverless  bool
	simulatedexecuteurl string
}

var instance *LyridClient
var once sync.Once

func GetInstance() *LyridClient {
	once.Do(func() {
		instance = &LyridClient{lyridaccess: os.Getenv("LYRID_ACCESS"), lyridsecret: os.Getenv("LYRID_SECRET"), simulateserverless: false}
	})
	return instance
}

func (client *LyridClient) SimulateServerless(url string) {
	client.simulateserverless = true
	client.simulatedexecuteurl = url
}

func (client *LyridClient) DisableSimulate() {
	client.simulateserverless = false
	client.simulatedexecuteurl = ""
}

func (client *LyridClient) GetLyridURL() string {
	if client.lyridurl == "" {
		client.lyridurl = "https://api.lyrid.io"
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

	// use
	client.lyridaccess = access
	client.lyridsecret = secret

	err := client.GetLyraVersion()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = client.login()
	if err != nil {
		log.Println(err)
		return err
	}
	//validatetoken()
	return nil
}

func (lc *LyridClient) GetLyraVersion() error {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: ""}
	response, err := cli.Get("/version")
	if err == nil {
		if response.StatusCode != 200 {
			var user model.LyridUser
			databyte, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(databyte, &user)
			return errors.New(string(databyte))
		}
	} else {
		return err
	}
	return nil

}

func (lc *LyridClient) GetUserProfile() *model.LyridUser {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	if lc.checktoken() {
		response, err := cli.Get("/api/user")
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
		response, err := cli.Get("/api/accounts")
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
	fmt.Println("Token: ", lc.token)
	if lc.checktoken() {
		response, err := cli.Get("/api/serverless/app/get")
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

func (lc *LyridClient) GetPublishedApps() []*model.PublishedApp {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}
	fmt.Println("Token:", lc.token)
	if lc.checktoken() {
		response, err := cli.Get("/api/serverless/install/")
		if err == nil {
			if response.StatusCode == 200 {
				var apps []*model.PublishedApp
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
		response, err := cli.Get("/api/serverless/app/get/" + AppId)
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
		response, err := cli.Get("/api/serverless/app/get/" + AppId + "/" + ModuleId)
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
		response, err := cli.Get("/api/serverless/app/get/" + AppId + "/" + ModuleId + "/" + RevisionId)
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

func (lc *LyridClient) ExecuteFunction(FunctionId string, Framework string, Body string) ([]byte, error) {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token}

	if lc.simulateserverless {
		cli.LyraUrl = lc.simulatedexecuteurl
	}

	if lc.checktoken() {
		response, err := cli.Post(lc.geturl("/api/serverless/app/execute/"+FunctionId+"/"+Framework), Body)
		if err == nil {
			if response.StatusCode == 200 {
				defer response.Body.Close()
				return ioutil.ReadAll(response.Body)
			}
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Unable to execute function.")
}

// POST /execute/:appname/:modulename/:tag/:functionname
func (lc *LyridClient) ExecuteFunctionByName(AppName string, ModuleName string, Tag string, FunctionName string, Body string) ([]byte, error) {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token, Access: model.UserAccessToken{Key: lc.lyridaccess, Secret: lc.lyridsecret}}

	if lc.simulateserverless {
		cli.LyraUrl = lc.simulatedexecuteurl
	}

	if lc.checktoken() {
		response, err := cli.PostBasicAuth(lc.geturl("/x/"+AppName+"/"+ModuleName+"/"+Tag+"/"+FunctionName+"/"), Body)
		if err == nil {
			if response.StatusCode == 200 {
				defer response.Body.Close()
				return ioutil.ReadAll(response.Body)
			}
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Unable to execute function.")
}

func (lc *LyridClient) ExecuteApp(AppName string, ModuleName string, Tag string, FunctionName string, Uri string, Method string, Body string) ([]byte, error) {
	cli := client.HTTPClient{LyraUrl: lc.GetLyridURL(), Token: lc.token, Access: model.UserAccessToken{Key: lc.lyridaccess, Secret: lc.lyridsecret}}

	if lc.simulateserverless {
		cli.LyraUrl = lc.simulatedexecuteurl
	}

	if lc.checktoken() {
		path := lc.geturl("/x/"+AppName+"/"+ModuleName+"/"+Tag+"/"+FunctionName) + Uri

		if Method == "GET" {
			response, err := cli.Get(path)
			if err == nil {
				if response.StatusCode == 200 {
					defer response.Body.Close()
					return ioutil.ReadAll(response.Body)
				}
			} else {
				return nil, err
			}
		} else if Method == "POST" {
			response, err := cli.Post(path, Body)
			if err == nil {
				if response.StatusCode == 200 {
					defer response.Body.Close()
					return ioutil.ReadAll(response.Body)
				}
			} else {
				return nil, err
			}
		} else if Method == "DELETE" {
			response, err := cli.Delete(path)
			if err == nil {
				if response.StatusCode == 200 {
					defer response.Body.Close()
					return ioutil.ReadAll(response.Body)
				}
			} else {
				return nil, err
			}
		}
	}
	return nil, errors.New("Unable to execute function.")
}

func (lc *LyridClient) geturl(path string) string {
	if lc.simulateserverless {
		return ""
	} else {
		return path
	}
}

func (lc *LyridClient) GetAccountPolicies() []*model.Policy {
	return nil
}

func (lc *LyridClient) GetModulePolicies(ModuleId string) []*model.Policy {
	return nil
}

func (client *LyridClient) checktoken() bool {
	if client.simulateserverless {
		return true
	}

	if client.istokenexpired() {
		client.login()
		return client.istokenexpired()
	}

	return true
}

func (client *LyridClient) istokenexpired() bool {
	if client.token == "" {
		log.Println("no token assigned")
		return true
	}

	token, _, err := new(jwt.Parser).ParseUnverified(client.token, jwt.MapClaims{})
	if err != nil {
		log.Println(err)
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		err = claims.Valid()
		if err != nil {
			log.Println(err)
		} else {
			// just renew token every 6 hours
			var tm time.Time
			switch iat := claims["iat"].(type) {
			case float64:
				tm = time.Unix(int64(iat), 0)
			case json.Number:
				v, _ := iat.Int64()
				tm = time.Unix(v, 0)
			}

			if tm.Add(time.Hour * 6).Before(time.Now()) {
				return true
			}

			return false
		}
	} else {
		log.Println(err)
	}

	return false
}

func (client *LyridClient) login() (string, error) {
	if client.simulateserverless {
		return "token_string", nil
	}

	jsonData := map[string]string{"key": client.lyridaccess, "secret": client.lyridsecret}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(client.GetLyridURL()+"/auth", "application/json", bytes.NewBuffer(jsonValue))
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
