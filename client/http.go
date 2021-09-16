package client

import (
	"bytes"
	"github.com/LyridInc/go-sdk/model"
	"net/http"
)

type HTTPClient struct {
	LyraUrl string
	Token   string
	Access  model.UserAccessToken
}

func (client *HTTPClient) Get(uri string) (*http.Response, error) {
	httpclient := &http.Client{}
	req, err := http.NewRequest("GET", client.LyraUrl+uri, nil)
	if err != nil {
		return nil, err
	}
	if client.Token != "" {
		req.Header.Add("Authorization", "Bearer "+client.Token)
	}
	return httpclient.Do(req)
}

func (client *HTTPClient) Delete(uri string) (*http.Response, error) {
	httpclient := &http.Client{}
	req, err := http.NewRequest("DELETE", client.LyraUrl+uri, nil)
	if err != nil {
		return nil, err
	}
	if client.Token != "" {
		req.Header.Add("Authorization", "Bearer "+client.Token)
	}
	return httpclient.Do(req)
}

func (client *HTTPClient) Post(uri string, body string) (*http.Response, error) {
	httpclient := &http.Client{}
	req, err := http.NewRequest("POST", client.LyraUrl+uri, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+client.Token)
	return httpclient.Do(req)
}

func (client *HTTPClient) PostBasicAuth(uri string, body string) (*http.Response, error) {
	httpclient := &http.Client{}
	req, err := http.NewRequest("POST", client.LyraUrl+uri, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.Access.Key, client.Access.Secret)
	return httpclient.Do(req)
}
