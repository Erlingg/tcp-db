package store

import "interface/domain"

type Mock struct {
}

func (m Mock) Add(u domain.User) {
}

func (m Mock) GetList() map[string]domain.User {
	return nil
}
