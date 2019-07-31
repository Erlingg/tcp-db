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
	case "GET":
		res = h.srv.GET(split[1])
	case "SET":
		res = h.srv.SET(split[1], split[2])
	case "DEL":
		res = h.srv.DEL(split[1])
	case "KEYS":
		res = h.srv.KEYS()
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
	return res + "\n"
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
