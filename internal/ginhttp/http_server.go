package ginhttp

import (
	"context"
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/service"
)

type Server struct {
	router  *gin.Engine
	service *service.Service
}

// NewServer ...
func NewServer(router *gin.Engine, service *service.Service) *Server {
	return &Server{
		router:  router,
		service: service,
	}
}

func (s *Server) Run(ctx context.Context) {

	// e.HTTPErrorHandler = handleEchoError(cfg)
	s.router.POST("/", s.Test)

	// runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	//runtimeCfg.BuildInfo = service.Version
	//runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	// routes
	// e.GET("/students", s.getStudents())
	//e.GET("/students/:studentId", s.getStudentById())
	//e.POST("/students", s.createStudent())
	//e.PUT("/students/:studentId", s.updateStudent())
	//e.DELETE("/students/:studentId", s.deleteStudent())

	// run actual server
	// echokit.RunServerWithContext(ctx, e, runtimeCfg)

	err := s.router.Run(":9070")
	if err != nil {
		panic(err)
	}
}
