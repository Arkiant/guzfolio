package datastore

import (
	"guzfolio/model"
)

type DataStore interface {
	GetUserByID(id uint) (*model.User, error)
	GetUserByIDs(ids []uint) ([]*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetCurrencyByID (id uint) (*model.Currency, error)
	GetCurrencyByIDs (ids []uint) ([]*model.Currency, error)
	GetAllCurrencies() ([]*model.Currency, error)
	GetAllUsers() ([]*model.User, error)
	GetPortfoliosByUserID(id uint) ([]*model.Portfolio, error)
	GetPortfoliosByUserIDs(id []uint) ([]*model.Portfolio, error)
	RegisterUser(input model.RegisterInput) (*model.User, error)
	CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error)
	CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error)
}