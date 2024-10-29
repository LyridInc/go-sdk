package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type (
	SentryIssue struct {
		ID            string `json:"id"`
		Title         string
		Annotations   []interface{}
		AssignedTo    string
		Count         string
		Filtered      map[string]interface{}
		FirstSeen     *time.Time
		HasSeen       bool
		IsBookmarked  bool
		IsPublic      bool
		IsSubscribed  bool
		IsUnhandled   bool
		IssueCategory string
	}
)

type SentryClient struct {
	DSN        string `json:"dsn"`
	APIBaseURL string `json:"apiBaseUrl"`
	AuthToken  string `json:"authToken"`
}

func parseApiURLFromDSN(dsn string) (string, error) {
	sentryDsnSplit := strings.Split(dsn, "@")
	if len(sentryDsnSplit) <= 1 {
		return "", fmt.Errorf("sentry dsn format is invalid")
	}
	sentryUrlSplit := strings.Split(sentryDsnSplit[1], "/")
	if len(sentryUrlSplit) <= 1 {
		return "", fmt.Errorf("sentry url format is invalid")
	}

	sentryUrl := "https://" + sentryUrlSplit[0] + "/api"
	return sentryUrl, nil
}

func NewSentryClient(dsn, authToken string) *SentryClient {
	apiBaseUrl, _ := parseApiURLFromDSN(dsn)
	return &SentryClient{
		DSN:        dsn,
		APIBaseURL: apiBaseUrl,
		AuthToken:  authToken,
	}
}

func (c *SentryClient) DoHttpRequest(method, url string, request interface{}) (*http.Response, error) {
	client := &http.Client{}
	var (
		req *http.Request
		err error
	)

	if request == nil {
		req, err = http.NewRequest(method, c.APIBaseURL+url, nil)
	} else {
		data, _ := json.Marshal(request)
		req, err = http.NewRequest(method, c.APIBaseURL+url, bytes.NewReader(data))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

func (c *SentryClient) GetIssuesList() (*http.Response, error) {
	return c.DoHttpRequest("GET", "/0/organizations/lyrid/issues/", nil)
}

func (c *SentryClient) GetIssueByTransactionID(transactionId string) (*http.Response, error) {
	return c.DoHttpRequest("GET", "/0/organizations/lyrid/issues/?query=transaction_id%3A"+transactionId, nil)
}

func (c *SentryClient) ResolveIssue(issueId string) (*http.Response, error) {
	data := struct {
		Status string `json:"status"`
	}{
		Status: "resolved",
	}
	return c.DoHttpRequest("PUT", "/0/organizations/lyrid/issues/"+issueId+"/", data)
}
