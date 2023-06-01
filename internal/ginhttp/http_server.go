package ginhttp

import (
	"context"
	"inventory-app-be/internal/service"

	"github.com/gin-gonic/gin"
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
	s.router.POST("/users", s.CreateUser)
	s.router.GET("/users", s.GetUserByID)

	s.router.GET("/products", s.getProducts)
	s.router.POST("/products", s.createProduct)
	s.router.GET("/products/:productID", s.getProductByID)
	s.router.PUT("/products/:productID", s.updateProductByID)
	s.router.DELETE("/products/:productID", s.deleteProductByID)
	// routes
	// e.GET("/students", s.getStudents())
	// e.GET("/students/:studentId", s.getStudentById())
	// e.POST("/students", s.createStudent())
	// e.PUT("/students/:studentId", s.updateStudent())
	// e.DELETE("/students/:studentId", s.deleteStudent())

	// run actual server
	// echokit.RunServerWithContext(ctx, e, runtimeCfg)

	err := s.router.Run(":9070")
	if err != nil {
		panic(err)
	}
}
