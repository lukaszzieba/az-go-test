package user

import (
	"context"

	"github.com/lukaszzieba/go-fun/internal/db"
)

type UserRepo interface {
	// Define methods for user storage, e.g., CreateUser, GetUser, etc.
	CreateUser(ctx context.Context, email string) (db.User, error)
	GetAllUsers(ctx context.Context) ([]db.User, error)
	DeleteAllUsers(ctx context.Context) error
}

type UserRepoSql struct {
	queries *db.Queries
}

func NewUserStorageSqlc(queries *db.Queries) *UserRepoSql {
	return &UserRepoSql{queries: queries}
}

func (s *UserRepoSql) CreateUser(ctx context.Context, email string) (db.User, error) {
	return s.queries.CreateUser(ctx, email)
}

func (s *UserRepoSql) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return s.queries.GetAllUsers(ctx)
}

func (s *UserRepoSql) DeleteAllUsers(ctx context.Context) error {
	return s.queries.DeleteAllUsers(ctx)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, email string) (db.User, error) {
	return s.repo.CreateUser(ctx, email)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return s.repo.GetAllUsers(ctx)
}
