package client

import "github.com/LyridInc/go-sdk/model"

type HTTPClient struct {
	LyraUrl string
	Token   string
	Access  model.UserAccessToken
}

func (client *HTTPClient) Initialize(token string) {

}

func (client *HTTPClient) Get(uri string) ([]byte, error) {

}

func (client *HTTPClient) Post(uri string, body string) ([]byte, error) {

}
