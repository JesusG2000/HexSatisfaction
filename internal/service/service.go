package service

import (
	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
)

// User is an interface for UserService methods.
type User interface {
	Create(req model.RegisterUserRequest) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(req model.LoginUserRequest) (string, error)
	IsExist(login string) (bool, error)
}

// UserRole is an interface for UserRoleService methods.
type UserRole interface {
	FindAllUser() ([]model.User, error)
}

// Purchase is an interface for PurchaseService repository methods.
type Purchase interface {
	Create(request model.CreatePurchaseRequest) (int, error)
	Delete(request model.DeletePurchaseRequest) (int, error)
	FindByID(request model.IDPurchaseRequest) (*model.Purchase, error)
	FindLastByUserID(request model.UserIDPurchaseRequest) (*model.Purchase, error)
	FindAllByUserID(request model.UserIDPurchaseRequest) ([]model.Purchase, error)
	FindByUserIDAndPeriod(request model.UserIDPeriodPurchaseRequest) ([]model.Purchase, error)
	FindByUserIDAfterDate(request model.UserIDAfterDatePurchaseRequest) ([]model.Purchase, error)
	FindByUserIDBeforeDate(request model.UserIDBeforeDatePurchaseRequest) ([]model.Purchase, error)
	FindByUserIDAndFileName(request model.UserIDFileNamePurchaseRequest) ([]model.Purchase, error)
	FindLast() (*model.Purchase, error)
	FindAll() ([]model.Purchase, error)
	FindByPeriod(request model.PeriodPurchaseRequest) ([]model.Purchase, error)
	FindAfterDate(request model.AfterDatePurchaseRequest) ([]model.Purchase, error)
	FindBeforeDate(request model.BeforeDatePurchaseRequest) ([]model.Purchase, error)
	FindByFileName(request model.FileNamePurchaseRequest) ([]model.Purchase, error)
}

// Services collects all service interfaces.
type Services struct {
	User     User
	UserRole UserRole
	Purchase Purchase
}

// Deps represents dependencies for services.
type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager
}

// NewServices is a Services constructor.
func NewServices(deps Deps) *Services {
	return &Services{
		User:     NewUserService(deps.Repos.User, deps.TokenManager),
		UserRole: NewUserRoleService(deps.Repos.UserRole),
		Purchase: NewPurchaseService(deps.Repos.Purchase),
	}
}
