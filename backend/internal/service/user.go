package service

import (
	"context"
	"errors"
	"github.com/mymikasa/mbook/backend/internal/domain"
	"github.com/mymikasa/mbook/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

//var ErrEmailAlreadyExisted = errors.New("邮箱已存在")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	//
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, err
	}

	//
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Edit(ctx context.Context, u domain.User) error {

	user, err := svc.repo.FindByEmail(ctx, u.Email)

	if !errors.Is(err, repository.ErrUserNotFound) {
		return ErrUserDuplicateEmail
	}

	user.Birthday = u.Birthday
	user.NickName = u.NickName

	// 更新
	return svc.repo.Update(ctx, u)
}
