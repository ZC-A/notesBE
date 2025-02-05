package http

import (
	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/ZC-A/notesBE/pkg/trace"
	"github.com/gin-gonic/gin"
)

func helloWorldHandler(c *gin.Context) {

	var err error

	ctx, span := trace.NewSpan(c, "helloWorldHandler => hello world")
	defer span.End(&err)
	log.Infof(ctx, "helloWorldHandler => hello world")
	c.JSON(200, gin.H{})
}
