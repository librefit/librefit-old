package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func offSearch(c *gin.Context) {
	offParams := "search_simple=1&action=process&json=1"
	country := c.DefaultQuery("country", "uk")

	offUrl := fmt.Sprintf("https://%s.openfoodfacts.org", country)

	for i, param := range c.Request.URL.Query() {
		offParams = fmt.Sprintf("%s&%s=%s", offParams, i, param[0])
	}

	remote, err := url.Parse(offUrl)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = "/cgi/search.pl"
		req.URL.ForceQuery = true
		req.URL.RawQuery = offParams
	}

	proxy.ModifyResponse = UpdateResponse

	proxy.ServeHTTP(c.Writer, c.Request)
}

func UpdateResponse(r *http.Response) error {
	// Origin comes with Access-Control-Allow-Origin: * but my cors also have them.
	// causing duplication. Browser doesn't like
	r.Header.Del("Access-Control-Allow-Origin")
	return nil
}
