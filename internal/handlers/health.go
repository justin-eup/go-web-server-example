package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler checks the health of the server
func HealthCheckHandler(c *gin.Context) {
	healthy := logFunc(isHealthy)

	var resp healhCheckResponse
	resp = &simpleHealthCheckResponse{isHealthy: healthy()}
	resp = &suffixResponse{response: resp}
	c.JSON(http.StatusOK, gin.H{
		"status": resp.getResponse(),
	})
}

func isHealthy() bool {
	return true
}

// decorate - func
type checkHealth func() bool

func logFunc(fn checkHealth) checkHealth {
	return func() bool {
		log.Println("start health checking...")
		isHealthy := fn()
		log.Println("finish health checking...")
		return isHealthy
	}
}

// decorate - struct
type healhCheckResponse interface {
	getResponse() string
}

type simpleHealthCheckResponse struct {
	isHealthy bool
}

func (s *simpleHealthCheckResponse) getResponse() string {
	if s.isHealthy {
		return "UP"
	}
	return "DOWN"
}

type suffixResponse struct {
	response healhCheckResponse
}

func (s *suffixResponse) getResponse() string {
	return s.response.getResponse() + "!!!"
}
