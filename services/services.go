package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/K31NER/go-gateway-fastapi/middleware"
	"github.com/K31NER/go-gateway-fastapi/schemas"
	"github.com/gin-gonic/gin"
)

var (
	users = make(map[int]schemas.Users)
    mu sync.Mutex
)

// Definimos el servidor de fastapi
func HandlerFastapi(ctx *gin.Context){
	fastapiURL := os.Getenv("FASTAPI_URL")
	if fastapiURL == "" {
		fastapiURL = "http://localhost:8000"
	}
	target,_ := url.Parse(fastapiURL)

	// Definimos el proxy inverso
	proxy := httputil.NewSingleHostReverseProxy(target)
	fastapi := middleware.MiddlewareFastapi(proxy)
	fastapi.ServeHTTP(ctx.Writer, ctx.Request)
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
    
	// Buscamos el usuario
	mu.Lock()
	user, exists := users[id]
    mu.Unlock()

	if !exists{
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"user":user,
	})
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
    mu.Lock() // Bloqueamos para asegurar que la generación de ID e inserción sean atómicas
	newUser.AddId() 
	users[newUser.Id] = newUser
	mu.Unlock() // Desbloqueamos

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Usuario creado con exito",
		"users":newUser,
	})
}

// elimina usuario
func DeleteUserById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido, debe ser entero"})
		return
	}
    
	// eliminamos
	mu.Lock()
	_,exists := users[id]
	if exists{
		delete(users, id)
	}
	mu.Unlock()

	// Validamos si existe
	if !exists{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
    
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Usuario eliminado con exito",
	})
}

// Actualiza el perfil de usuario
func UpdateHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido, debe ser entero"})
		return
	}

	// Buscamos si existe el usuario
	mu.Lock()
	_, exists := users[id]
	mu.Unlock()

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var updatedUser schemas.Users
	// Decodificamos el JSON enviado en la solicitud
	if err := ctx.BindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar json"})
		return
	}

	// Aseguramos que el ID se mantenga igual
	updatedUser.Id = id

	// Actualizamos el usuario en el mapa
	mu.Lock()
	users[id] = updatedUser
	mu.Unlock()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuario actualizado con exito",
		"user":    updatedUser,
	})
}



