package repository

import (
	"context"
	"github.com/mymikasa/mbook/backend/internal/domain"
	"github.com/mymikasa/mbook/backend/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func (r *UserRepository) FindById(int64) {
	// 先从cache里找
	// 再从dao里找
	//

}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
