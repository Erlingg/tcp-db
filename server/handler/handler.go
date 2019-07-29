package handler

import (
	"strings"
	"tcp-db/server/domain"
)

type dbHandlers struct {
	srv domain.DBService
}

func NewDBHandlers(s domain.DBService) *dbHandlers {
	return &dbHandlers{srv: s}
}

func (h dbHandlers) Query(query string) string {
	split := strings.Split(query, " ")
	var res string
	switch split[0] {
		case "POP":
			res = h.srv.POP(split[1])
		case "PUT":
			h.srv.PUT(split[1], split[2])
		case "DEL":
			h.srv.DEL(split[1])
		case "LIST":
			res = h.srv.LIST()
	}
	//list := h.srv.GetServiceList()
	//
	//b, err := json.Marshal(list)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//_, _ = w.Write(b)
	return res
}
//
//func (h dbHandlers) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	u := domain.User{}
//	err := json.NewDecoder(r.Body).Decode(&u)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	h.srv.AddService(u)
//}
//
//func (h dbHandlers) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
//func (h dbHandlers) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
