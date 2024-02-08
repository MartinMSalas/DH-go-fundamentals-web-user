package user

import (
	domain "DH-go-fundamentals-web-user/internal/domain"
	"context"
	"log"
)

type DB struct {
	Users []domain.User
	MaxUserD uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}
	repo struct {
		db DB
		log *log.Logger
	}
)

func NewRepo(db DB, l *log.Logger) Repository {
	return &repo{
		db: db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserD++
	user.ID = r.db.MaxUserD
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("repository created")
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("repository get all")
	return r.db.Users, nil
}
