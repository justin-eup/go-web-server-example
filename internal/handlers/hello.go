package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler handles requests to /api/hello/:name
func HelloHandler(c *gin.Context) {
	name := c.Param("name")

	var response *helloResponse
	response = newNormalHelloResponseBuilder().setTemplate("Hello, %s!").setName(name).build()

	res := fmt.Sprintf(response.template, response.name)
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

type helloResponse struct {
	template string
	name     string
}

// builder - interface
type helloResponseBuilder interface {
	setTemplate(string) helloResponseBuilder
	setName(string) helloResponseBuilder
	build() *helloResponse
}

// builder - concrete implementation
type normalHelloResponseBuilder struct {
	template string
	name     string
}

func newNormalHelloResponseBuilder() *normalHelloResponseBuilder {
	return &normalHelloResponseBuilder{}
}

func (b *normalHelloResponseBuilder) setTemplate(template string) helloResponseBuilder {
	b.template = template
	return b
}

func (b *normalHelloResponseBuilder) setName(name string) helloResponseBuilder {
	b.name = name
	return b
}

func (b *normalHelloResponseBuilder) build() *helloResponse {
	res := &helloResponse{b.template, b.name}
	return res
}
