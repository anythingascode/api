package main

import (
	"github.com/anythingascode/api/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	for k, v := range endpoints.GetEndPoints {
		router.GET(k, v)
	}
	for k, v := range endpoints.PostEndPoints {
		router.POST(k, v)
	}
	router.RunTLS("seapi:8080", "ssl/cert.pem", "ssl/key.pem")
}
