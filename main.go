package main

import (
	"log"

	"github.com/K31NER/go-gateway-fastapi/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargamos las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	// Inistnaciamos el servidor de gin
	r := gin.Default()
    
	// Agregamos los routers
    routes.SetUpRoutes(r)
    
	// Corremos el servidor
	r.Run(":8080")
}