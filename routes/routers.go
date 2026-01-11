package routes

import (
	"github.com/K31NER/go-gateway-fastapi/services"
	"github.com/gin-gonic/gin"
)

// Funcion que almacena los routes
func SetUpRoutes(r *gin.Engine){
    // Rutas
	r.GET("/health", services.HandlerHealth)
	r.GET("/users/:id", services.HandlerGetUserById)
	r.GET("/users", services.HandlerReadAllUsers)
    r.POST("/users",services.HandlerAddUser)

	// Si ninguna ruta coincide se usan las de fastapi
	r.NoRoute(services.HandlerFastapi)
}
