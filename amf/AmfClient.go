package amf

import (
	"net/http"
	
)

var (
	NsmfPduSessionClient *http.Client
	SmfScheme string
	SmfHost string
	SmfBasePath string
	SmfBaseURL string
	Boundary string
	AmfID string
)
func init(){
	NsmfPduSessionClient=&http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	SmfScheme = "http"
	SmfHost = "localhost:8080"
	SmfBasePath = "/nsmf-pdusession/v1"
	Boundary = "----Boundary"
	AmfID = "localhost:8080"
	SmfBaseURL = SmfScheme+"://"+SmfHost+SmfBasePath
}

// define customize redirect policy
// via: old request
// req: new request
func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
	if req==nil || via == nil{
		return nil
	}
	
	return nil
}

// 下面的數據用於認證

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")
	
	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")
	
	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")
	
	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}