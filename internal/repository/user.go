package repository

import (
	"onelab2/internal/common"
	"onelab2/internal/model"
	"sync"

	"go.uber.org/zap"
)

type UserRepository struct {
	mp  map[string]*model.User
	rw  sync.RWMutex
	log *zap.Logger
}

func NewUserRepository(log *zap.Logger) *UserRepository {
	return &UserRepository{
		mp:  make(map[string]*model.User),
		log: log,
	}
}

func (r *UserRepository) Create(u *model.User) error {
	r.rw.Lock()
	defer r.rw.Unlock()

	r.log.Info("Create user", zap.String("login", u.Login))
	r.mp[u.Login] = u

	return nil
}

func (r *UserRepository) GetByLogin(login string) (*model.User, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	r.log.Info("Get user by login", zap.String("login", login))
	u, ok := r.mp[login]
	if !ok {
		r.log.Info("User not found", zap.String("login", login))
		return nil, common.ErrUserNotFound
	}

	return u, nil
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	var users []*model.User

	r.log.Info("Get all users")

	for _, u := range r.mp {
		users = append(users, u)
	}

	return users, nil
}
