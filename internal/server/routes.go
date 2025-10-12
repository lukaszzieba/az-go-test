package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /users", s.users())
	mux.HandleFunc("/", s.handler())

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (server *Server) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := map[string]any{"success": true}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func (server *Server) users() http.HandlerFunc {
	type UserResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := server.userService.GetAllUsers(r.Context())
		fmt.Println(users)
		res := make([]UserResponse, 0)
		for _, u := range users {
			res = append(res, UserResponse{
				ID:    u.ID.String(),
				Email: u.Email,
			})
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
