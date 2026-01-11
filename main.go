package main

import (
	"github.com/K31NER/go-gateway-fastapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Inistnaciamos el servidor de gin
	r := gin.Default()
    
	// Agregamos los routers
    routes.SetUpRoutes(r)
    
	// Corremos el servidor
	r.Run(":8080")
}