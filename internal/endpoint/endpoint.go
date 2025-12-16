package endpoint

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lavatee/tracker_backend/internal/service"
)

type Endpoint struct {
	services *service.Service
}

func NewEndpoint(services *service.Service) *Endpoint {
	return &Endpoint{
		services: services,
	}
}

func (e *Endpoint) InitRoutes() *gin.Engine {
	router := gin.New()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", e.SignUp)
		auth.POST("/sign-in", e.SignIn)
		auth.POST("/refresh", e.Refresh)
	}
	api := router.Group("/api", e.Middleware)
	{
		api.PUT("/user-balance", e.UpdateUserBalance)

		api.GET("/users/:id", e.GetOneUser)
		api.GET("/referral-users", e.GetUserReferrals)
		api.GET("/next-nodes/:id", e.GetNextNodes)
		api.GET("/previous-nodes/:id", e.GetPreviousNodes)
		api.PUT("/nodes/:id", e.PutNode)
		api.POST("/nodes", e.PostNode)
		api.GET("/nodes/:id", e.GetOneNode)

		api.POST("/achievements", e.CreateAchievement)
		api.GET("/achievements/my", e.GetMyAchievements)
		api.GET("/achievements/:id", e.GetAchievementById)
		api.GET("/achievements", e.GetAchievementsByStatus) // статус указывается так /api/achievements?status=... (также необходимо слушать YASO)
		api.DELETE("/achievements/:id", e.DeleteAchievement)
		api.POST("/achievements/:id/approve", e.ApproveAchievement)
		api.POST("/achievements/:id/reject", e.RejectAchievement)

		api.POST("/products", e.CreateProduct)
		api.GET("/products", e.GetProducts)
		api.GET("/products/:id", e.GetProductById)
		api.DELETE("/products/:id", e.DeleteProduct)
		api.PUT("/products/:id", e.UpdateProduct)

		api.POST("/cart", e.AddProductToCart)
		api.GET("/cart", e.GetUserCart)
		api.PUT("/cart/:productId", e.UpdateProductInCartAmount)
		api.DELETE("/cart/:productId", e.DeleteProductFromCart)
		api.DELETE("/cart", e.CleanUserCart)

		api.POST("/orders", e.CreateOrder)
		api.GET("/orders/my", e.GetMyOrders)
		api.GET("/orders/:id", e.GetOrderById)
		api.GET("/orders", e.GetOrdersByStatus) // статус указывается так /api/orders?status=...
		api.POST("/orders/:id/reject", e.SetRejectedStatus)
		api.POST("/orders/:id/ready", e.SetReadyStatus)
		api.POST("/orders/:id/issue", e.SetIssuedStatus)
	}
	return router
}
