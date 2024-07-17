package api

import (
	handler "api_getway/api/handlers"
	"api_getway/api/middlewares"

	_ "api_getway/api/docs"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// @title EcoSwap System API
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type Server struct {
	Handlers handler.Handlers
}

func NewServer(hands handler.Handlers) *Server {
	return &Server{Handlers: hands}
}

func (s *Server) InitRoutes(r *gin.Engine) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	
	auth := s.Handlers.Auth()
	item := s.Handlers.Item()
	
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middlewares.JWTMiddlewares)


	authGroup := r.Group("/auth")
	{
		authGroup.GET("/profile", auth.GetProfile)
		authGroup.PUT("/profile", auth.EditProfile)
		authGroup.GET("/users", auth.ListUsers)
		authGroup.DELETE("/users/:id", auth.DeleteUser)
		authGroup.POST("/reset-password", auth.ResetPassword)
		authGroup.POST("/refresh", auth.RefreshToken)
		authGroup.POST("/logout", auth.Logout)
		authGroup.GET("/eco-points", auth.GetEcoPoints)
		authGroup.POST("/eco-points", auth.AddEcoPoints)
		authGroup.GET("/eco-points/history", auth.GetEcoPointsHistory)
	}

	itemGroup := r.Group("/items")
	{
		itemGroup.POST("/", item.CreateItem)
		itemGroup.PUT("/:id", item.UpdateItem)
		itemGroup.DELETE("/:id", item.DeleteItem)
		itemGroup.GET("/:id", item.GetItem)
		itemGroup.GET("/", item.GetItems)
		itemGroup.GET("/search", item.SearchItem)
	}

	swapGroup := r.Group("/swaps")
	{
		swapGroup.POST("/", item.CreateSwap)
		swapGroup.PUT("/:id/accept", item.UpdateSwap)
		swapGroup.GET("/", item.GetSwaps)
	}

	recyclingGroup := r.Group("/recycling-centers")
	{
		recyclingGroup.POST("/", item.AddRecyclingCentr)
		recyclingGroup.GET("/", item.GetRecyclingCentrs)
		recyclingGroup.POST("/submissions", item.AddRecyclingSubmission)
	}

	ratingGroup := r.Group("/users/:id/ratings")
	{
		ratingGroup.POST("/", item.AddUserRating)
		ratingGroup.GET("/", item.GetUserRatings)
	}

	itemCategoryGroup := r.Group("/item-categories")
	{
		itemCategoryGroup.POST("/", item.AddItemCategory)
	}

	statisticsGroup := r.Group("/statistics")
	{
		statisticsGroup.GET("/", item.GetStatistics)
	}

	userActivityGroup := r.Group("/user-activity/:id")
	{
		userActivityGroup.GET("/", item.GetUserActivity)
	}

	ecoChallengeGroup := r.Group("/eco-challenges")
	{
		ecoChallengeGroup.POST("/", item.AddEcoChallenge)
		ecoChallengeGroup.POST("/:id/participate", item.ParticipateEcoChallenge)
		ecoChallengeGroup.PUT("/:id/update-progress", item.UpdateEcoChallengeProgress)
	}

	ecoTipGroup := r.Group("/eco-tips")
	{
		ecoTipGroup.POST("/", item.AddEcoTip)
		ecoTipGroup.GET("/", item.GetEcoTips)
	}

}
