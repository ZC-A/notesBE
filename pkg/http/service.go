package http

import (
	"context"
	gohttp "net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ZC-A/notesBE/pkg/log"
	"github.com/gin-gonic/gin"
)

// Service
type Service struct {
	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc

	// 全局唯一的http服务
	server *gohttp.Server
	g      *gin.Engine
}

// Type
func (s *Service) Type() string {
	return "http"
}

// Start
func (s *Service) Start(ctx context.Context) {
	s.Reload(ctx)
}

// Reload
func (s *Service) Reload(ctx context.Context) {
	var err error

	// 先关闭当前的服务
	if s.server != nil {
		log.Warnf(ctx, "http server is running, will stop it first, max waiting time->[%s].", WriteTimeout)
		tempCtx, cancelFunc := context.WithTimeout(ctx, WriteTimeout)
		defer cancelFunc()
		if err = s.server.Shutdown(tempCtx); err != nil {
			log.Errorf(ctx, "shutdown server with err->[%s]", err)
		}
		log.Warnf(ctx, "http server shutdown done.")
	}

	if s.cancelFunc != nil {
		s.cancelFunc()
	}

	log.Debugf(ctx, "waiting for http service close")
	s.Wait()

	gin.SetMode(gin.ReleaseMode)
	s.g = gin.New()

	public := s.g.Group("/")
	// 注册默认路由
	// 注册中间件，注意中间件必须要在其他服务之前注册，否则中间件不生效
	public.Use(
		gin.Recovery(),
	)
	registerDefaultHandlers(ctx, public)

	// 构造新的http服务
	s.server = &gohttp.Server{
		Addr:         strings.Join([]string{IPAddress, strconv.Itoa(Port)}, ":"),
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		Handler:      s.g,
	}

	s.wg.Add(1)
	go func(server *gohttp.Server) {
		defer s.wg.Done()
		if err = server.ListenAndServe(); err != nil && err != gohttp.ErrServerClosed {
			log.Panicf(ctx, "failed to start server for->[%s]", err)
			return
		}
		log.Warnf(ctx, "last http server is closed now")
	}(s.server)

	s.ctx, s.cancelFunc = context.WithCancel(ctx)
	log.Debugf(ctx, "http service context update success.")
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		<-s.ctx.Done()
		err = s.server.Close()
		if err != nil {
			log.Errorf(ctx, "get error when closing http server:%s", err)
		}
	}()
	log.Infof(ctx, "http service reloaded or start success.")
}

// Wait
func (s *Service) Wait() {
	s.wg.Wait()
}

// Close
func (s *Service) Close() {
	s.cancelFunc()
	log.Infof(s.ctx, "http service context cancel func called.")
}
