package http

import (
	"context"
	"net/http"

	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterHandlers struct {
	ctx context.Context
	g   *gin.RouterGroup
}

func (r *RegisterHandlers) register(method, handlerPath string, handlerFunc ...gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		r.g.GET(handlerPath, handlerFunc...)
	case http.MethodPost:
		r.g.POST(handlerPath, handlerFunc...)
	case http.MethodHead:
		r.g.HEAD(handlerPath, handlerFunc...)
	default:
		log.Errorf(r.ctx, "registerHandlers error type is error %s", method)
		return
	}

	log.Infof(r.ctx, "registerHandlers => [%s] %s", method, handlerPath)
}

func getRegisterHandlers(ctx context.Context, g *gin.RouterGroup) *RegisterHandlers {
	return &RegisterHandlers{
		ctx: ctx,
		g:   g,
	}
}

func registerDefaultHandlers(ctx context.Context, g *gin.RouterGroup) {
	// TODO: register default handlers
	var handlerPath string

	registerHandler := getRegisterHandlers(ctx, g)

	handlerPath = viper.GetString(HelloWorldPathConfigPath)
	registerHandler.register(http.MethodGet, handlerPath, helloWorldHandler)
}
