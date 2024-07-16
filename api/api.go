package api

import (
	_ "bankomat/api/docs"
	"bankomat/api/handler"
	"bankomat/config"
	"bankomat/storage"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.POST("/accounts", handler.CreateAccount)
	r.POST("/accounts/:id/deposit", handler.Deposit)
	r.POST("/accounts/:id/withdraw", handler.Withdraw)
	r.GET("/accounts/:id/balance", handler.GetBalance)

	r.Use(customCORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
