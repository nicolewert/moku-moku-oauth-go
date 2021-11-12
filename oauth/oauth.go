package oauth

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var oauthRestClient = rest.RequestBuilder{
	BaseURL: "http://localhost:8080",
	Timeout: 200 * time.Millisecond,
}

type accessToken struct {
	Id     string `json:"id"`
	UserId int64  `json:"user_id"`
}
