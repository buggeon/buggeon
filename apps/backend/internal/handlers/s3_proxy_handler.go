package handlers

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func S3ProxyHandler(endpoint string) gin.HandlerFunc {

	target, _ := url.Parse(endpoint)
	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/files")
		proxy.ServeHTTP(c.Writer, c.Request)
	}

}
