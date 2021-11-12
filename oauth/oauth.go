package oauth

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var oauthRestClient = rest.RequestBuilder{
	BaseURL: "http://localhost:8080",
	Timeout: 200 * time.Millisecond,
}
