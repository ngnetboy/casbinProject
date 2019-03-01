package service

import (
	"model"
	"sync"
)

var Account = &accountService{
	mutex: &sync.Mutex{},
}

type accountService struct {
	mutex *sync.Mutex
}

func (a *accountService) GetAccountByName(name string) *model.Account {
	result := &model.Account{}
	if err := db.Where("name = ?", name).First(result).Error; err != nil {
		return nil
	}
	return result
}
