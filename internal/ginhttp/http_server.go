package ginhttp

import (
	"context"
	"inventory-app-be/internal/middleware"
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
	s.router.Use(middleware.CORSMiddleware())

	// s.router.POST("/forget-password", s.login)
	// s.router.POST("/reset-password", s.login)
	s.router.POST("/login", s.login)

	s.router.Use(middleware.Auth())

	userGroup := s.router.Group("/users")
	{
		userGroup.POST("/", s.createUser)
		// userGroup.PUT("/profile", s.createUser)
		userGroup.POST("/change-password", s.updateUserPasswordByID)
		userGroup.GET("/details", s.getUserDetailByToken)
		userGroup.PUT("/:userID", s.updateUserByID)
		userGroup.DELETE("/:userID", s.deleteUserByID)
	}

	// admin
	s.router.POST("/admins/registers", s.createUser)

	// customers
	s.router.POST("/customers", s.createCustomer)
	s.router.GET("/customers", s.getCustomers)
	s.router.GET("/customers/:customerID", s.getCustomerByID)
	s.router.PUT("/customers/:customerID", s.updateCustomerByID)
	s.router.DELETE("/customers/:customerID", s.deleteCustomerByID)

	// transactions customer
	s.router.POST("/transactions", s.createTransaction)
	s.router.GET("/transactions", s.getTransactions)
	s.router.POST("/transactions/status", s.getTransactionCustomersByStatus)
	s.router.POST("/transactions/sends/:transactionID", s.updateStatusTransactionCustomerSedangDikirimByID)
	s.router.POST("/transactions/finish/:transactionID", s.updateStatusTransactionCustomerSelesaiByID)
	s.router.GET("/transactions/:transactionID", s.getTransactionByID)
	s.router.PUT("/transactions/:transactionID", s.updateTransactionByID)
	s.router.DELETE("/transactions/:transactionID", s.deleteTransactionByID)

	// transactions supplier
	s.router.POST("/transactions/suppliers", s.createTransactionSupplier)
	s.router.GET("/transactions/suppliers", s.getTransactionSuppliers)
	s.router.POST("/transactions/suppliers/status", s.getTransactionSupplierByStatus)
	s.router.POST("/transactions/suppliers/:transactionID", s.updateStatusTransactionByIDAndUpdateStock)
	// s.router.GET("/transactions/:transactionID", s.getTransactionByID)
	//s.router.PUT("/transactions/:transactionID", s.updateTransactionByID)
	//s.router.DELETE("/transactions/:transactionID", s.deleteTransactionByID)

	// suppliers
	s.router.POST("/suppliers", s.createSupplier)
	s.router.GET("/suppliers", s.getSuppliers)
	s.router.GET("/suppliers/:supplierID", s.getSupplierByID)
	s.router.PUT("/suppliers/:supplierID", s.updateSupplierByID)
	s.router.DELETE("/suppliers/:supplierID", s.deleteSupplierByID)

	// products
	s.router.GET("/products", s.getProducts)
	s.router.GET("/products/suppliers/:supplierID", s.getProductBySupplierID)
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
