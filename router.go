package main

import "github.com/gin-gonic/gin"

func ConfigRoutes(r *gin.Engine) {
	r.Use(CORSMiddleware())
	r.POST("v1/users", CreateUser)
	r.GET("v1/users", GetUsers)
	r.PUT("v1/users", UpdateUser)
	r.DELETE("v1/users", DeleteUser)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
