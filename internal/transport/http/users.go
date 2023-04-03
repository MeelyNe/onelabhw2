package routes

import (
	"net/http"
	"onelab2/internal/model"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type userRepo interface {
	Create(u *model.User) error
	GetByLogin(login string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type User struct {
	repo userRepo
	l    *zap.Logger
}

func NewUser(repo userRepo, l *zap.Logger) *User {
	return &User{repo: repo, l: l}
}

func (u *User) NewUser(e echo.Context) error {
	var user model.User
	if err := e.Bind(&user); err != nil {
		u.l.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := u.repo.Create(&user); err != nil {
		u.l.Error("Create error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	u.l.Info("User created", zap.String("login", user.Login))

	return e.JSON(http.StatusOK, user)
}

func (u *User) GetUser(e echo.Context) error {
	login := e.Param("login")

	u.l.Info("Get user", zap.String("login", login))

	user, err := u.repo.GetByLogin(login)
	if err != nil {
		u.l.Info("Get user error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}

	u.l.Info("User found", zap.String("login", user.Login))
	return e.JSON(http.StatusOK, user)
}

func (u *User) GetAllUsers(e echo.Context) error {
	u.l.Info("Get all users")
	users, err := u.repo.GetAll()
	if err != nil {
		u.l.Error("Get all users error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, err)
	}
	u.l.Info("All users found", zap.Int("count", len(users)))

	return e.JSON(http.StatusOK, users)
}

func (u *User) Register(e *echo.Echo) {
	e.POST("/users", u.NewUser)
	e.GET("/users/:login", u.GetUser)
	e.GET("/users", u.GetAllUsers)
}
