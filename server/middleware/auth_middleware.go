package middleware

import "net/http"

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(rw, r)
	})
}

// Implement the ServeHTTP method to satisfy http.Handler
func (middleware *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Call the enableCORS middleware and pass the next handler
	corsHandler := middleware.enableCORS(middleware.Handler)
	corsHandler.ServeHTTP(rw, r)
}
