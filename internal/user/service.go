package user
import (
	"context"
	"DH-go-fundamentals-web-user/internal/domain"
	"log"

)
type (
	Service interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User,error)

	}

	service struct	{
		log *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, r Repository) Service {
	return &service{
		log: l,
		repo: r,
	}
}

func (s *service) Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
	}
	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Println("service get all")
	return users, nil
}