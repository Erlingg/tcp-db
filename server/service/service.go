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


func (s dbService) POP(k string) string {
	return s.store.GetRow(k) + "\n"
}
func (s dbService) PUT(k string, v string) {
	s.store.AddRow(k, v)
}
func (s dbService) DEL(k string) {
	s.store.DelRow(k)
}
func (s dbService) LIST() string {
	list := s.store.GetList()
	var res string
	for _, v := range list{
		res += v + "\n"
	}
	return res
}