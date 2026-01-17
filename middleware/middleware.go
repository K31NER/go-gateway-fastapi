package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// Middleware que meneja el trafico entrante a fastapi
func MiddlewareFastapi(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Obtenemos el header
        apiKey := r.Header.Get("X-API-Key")
        
		// Validamos la api key
		if apiKey != os.Getenv("API_KEY") || apiKey == "" {
			w.Header().Set("Conetent-Type","application/json")
			w.WriteHeader(http.StatusUnauthorized)
            
			// Definimos el mensaje de error
			response := map[string]string{
				"error":"Acceso denegado",
				"message":"Valide sus credenciales de acceso",
			}

			// codificamos la respuesta
		    _ = json.NewEncoder(w).Encode(response)
			return 
		}

		// Logs de entrada
        log.Printf("Solicitud enviada a Fastapi: %s - %s",r.Method,r.URL.Path)
        start := time.Now()

		// Pasamos la solicitud
		next.ServeHTTP(w,r)

		duration := time.Since(start)
        
		// Logs finales
		log.Printf("Tiempo de solicitud: %d ms",duration.Microseconds())
	})
}