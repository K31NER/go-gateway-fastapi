package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"

	"github.com/K31NER/go-gateway-fastapi/schemas"
	"github.com/gin-gonic/gin"
)

var (
	users []schemas.Users
    mu sync.Mutex
)

// Definimos el servidor de fastapi
func HandlerFastapi(ctx *gin.Context){
	target,_ := url.Parse("http://localhost:8000")

	// Definimos el proxy inverso
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func HandlerHealth(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
			"message":"Server is running",
		})
}

// Metodos crud

// Lee todos los usuarios
func HandlerReadAllUsers(ctx *gin.Context){
    
	ctx.JSON(http.StatusOK,gin.H{
		"users":users,
	})
}

// Busca usuario por id
func HandlerGetUserById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido, debe ser entero"})
		return
	}

	for _, user := range users {
		if user.Id == id {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}

// Agrega nuevo usuario
func HandlerAddUser(ctx *gin.Context){

    var newUser schemas.Users
    
	// Validamos los campos
	if err := ctx.BindJSON(&newUser); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Error al decodificar json",
		})
		return
	}
    
	if newUser.Name == "" || newUser.Mail == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Error los campos nombre y correo son requeridos",
		})
		return
	}

	// Registramos el nuevo usuario
	newUser.AddId() // Llamamos a tu método para asignar ID

    mu.Lock() // Bloqueamos la variable users
	users = append(users, newUser)
	mu.Unlock() // Desbloqueamos

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Usuario creado con exito",
		"users":users,
	})
}





