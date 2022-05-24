package engines

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/urls"
)

type NetHttpEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
	Translation    *lang.Translations

	Response response.NetHttpResponse
	Request  request.NetHttpRequest
	URL      urls.NetHttpURL
}

func NetHttp(w http.ResponseWriter, r *http.Request, t *lang.Translations) *NetHttpEngine {
	res := response.NetHttp(w, r)
	url := urls.NetHttp(w, r)
	req := request.NetHttp(w, r, t, url)

	return &NetHttpEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Translation:    t,
		Response:       *res,
		Request:        *req,
		URL:            *url,
	}
}

func (m NetHttpEngine) GetRequest() interface{} {
	return m.Request
}

func (m NetHttpEngine) GetResponse() interface{} {
	return m.Response
}

func (m NetHttpEngine) GetURL() interface{} {
	return m.URL
}
