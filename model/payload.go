package model

type Subdomain struct {
	Name string `json:"name" binding:"required"`

	AppId        string `json:"appId" binding:"required"`
	ModuleId     string `json:"moduleId" binding:"required"`
	FunctionName string `json:"functionName" binding:"required"`
	Tag          string `json:"tag"`
}

type DeployedService struct {
	Name    string      `json:"name" binding:"required"`
	Payload interface{} `json:"payload" binding:"required"`
}

type RequestPayload struct {
	Headers               map[string]string `json:"headers"`
	Path                  string            `json:"path"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`

	HttpMethod     string                 `json:"httpMethod"`
	RequestContext map[string]interface{} `json:"requestContext"`
	Body           []byte                 `json:"body"`

	IsBase64Encoded bool `json:"isBase64Encoded"`
}

func (request *RequestPayload) ToQuery() string {
	if len(request.QueryStringParameters) == 0 {
		return ""
	}

	queryReturn := ""

	for key, value := range request.QueryStringParameters {
		if queryReturn != "" {
			queryReturn = queryReturn + "&"
		}
		queryReturn = queryReturn + key + "=" + value
	}

	return queryReturn
}

type ResponsePayload struct {
	Headers             map[string]string   `json:"headers"`
	MultiValueHeaders   map[string][]string `json:"multiValueHeaders"`
	StatusCode          int                 `json:"statusCode"`
	Body                string              `json:"body"`
	ExecutionDurationMs int64               `json:"executionDurationMs"`

	IsBase64Encoded bool `json:"isBase64Encoded"`
}
