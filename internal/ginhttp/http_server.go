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

	// products
	s.router.GET("/products", s.getProducts)
	s.router.POST("/products", s.createProduct)
	s.router.GET("/products/:productID", s.getProductByID)
	s.router.PUT("/products/:productID", s.updateProductByID)
	s.router.DELETE("/products/:productID", s.deleteProductByID)

	// product variants
	s.router.GET("/products/variants", s.getProductVariants)
	s.router.POST("/products/variants", s.createProductVariant)
	s.router.GET("/products/:productID/variants", s.getProductVariantByProductID)
	s.router.GET("/products/variants/:productVariantID", s.getProductVariantByID)
	s.router.PUT("/products/variants/:productVariantID", s.updateProductVariantByID)
	s.router.DELETE("/products/variants/:productVariantID", s.deleteProductVariantByID)

	err := s.router.Run(":9070")
	if err != nil {
		panic(err)
	}
}
