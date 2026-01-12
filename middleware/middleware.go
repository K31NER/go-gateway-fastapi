package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware que meneja el trafico entrante a fastapi
func MiddlewareFastapi(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
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