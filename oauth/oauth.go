package oauth

import (
	"net/http"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXPublic = "X-Public"
)

var oauthRestClient = rest.RequestBuilder{
	BaseURL: "http://localhost:8080",
	Timeout: 200 * time.Millisecond,
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}
