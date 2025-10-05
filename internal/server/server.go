package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lukaszzieba/az-go-test/internal/storage/user"
)

type Server struct {
	port        int
	userService *user.UesrService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbUrl := os.Getenv("DB_URL")

	dbQueries := NewSqlcQueries(dbUrl)
	userService := user.NewUserService(user.NewUserStorageSqlc(dbQueries))

	NewServer := &Server{
		port:        port,
		userService: userService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
