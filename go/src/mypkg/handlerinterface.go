package mypkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseHandler interface {
	Handle(c *gin.Context, outboundcon []*http.Response) (int, interface{})
}
