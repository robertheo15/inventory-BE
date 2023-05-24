package echohttp

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
	//studentService *student.StudentService
}

// NewServer ...
func NewServer(router *echo.Echo) *Server {
	return &Server{
		router: router,
		//studentService: studentService,
	}
}

func (s *Server) Run(ctx context.Context) {
	e := echo.New()
	//e.HTTPErrorHandler = handleEchoError(cfg)

	//runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	//runtimeCfg.BuildInfo = service.Version
	//runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	// routes
	//e.GET("/students", s.getStudents())
	//e.GET("/students/:studentId", s.getStudentById())
	//e.POST("/students", s.createStudent())
	//e.PUT("/students/:studentId", s.updateStudent())
	//e.DELETE("/students/:studentId", s.deleteStudent())

	// run actual server
	//echokit.RunServerWithContext(ctx, e, runtimeCfg)
	e.Logger.Fatal(e.Start(":9070"))

}
