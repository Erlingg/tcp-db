package service

import "tcp-db/server/domain"

type dbService struct {
	store domain.DBStore
}

func NewDBService(s domain.DBStore) *dbService {
	return &dbService{store: s}
}

//func (b dbService) AddService(u domain.User) {
//
//
//	b.store.AddStore(u)
//}
//
//func (b dbService) GetServiceList() map[string]domain.User {
//	return b.store.GetStoreList()
//}
//
//func (b dbService) DeleteService()  {
//
//}

func (s dbService) GET(k string) string {
	res, err := s.store.GetRow(k)
	if err != nil {
		return err.Error()
	}
	return res
}
func (s dbService) SET(k string, v string) string {
	err := s.store.SetRow(k, v)
	if err != nil {
		return err.Error()
	}
	return "SET succeeded"
}
func (s dbService) DEL(k string) string {
	err := s.store.DelRow(k)
	if err != nil {
		return err.Error()
	}
	return "DEL succeeded"
}
func (s dbService) KEYS() string {
	list := s.store.GetKeys()
	if len(list) == 0 {
		return "Empty storage"
	}
	var res string
	for _, v := range list {
		res += v + ","
	}
	return res[:len(res)-1]
}
