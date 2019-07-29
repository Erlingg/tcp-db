package domain

//type Table struct {
//	Value string
//}

type DBStore interface {
	GetRow(k string) string
	AddRow(k string, v string)
	DelRow(k string)
	GetList() map[string]string
}

type DBService interface {
	POP(k string) string
	PUT(k string, v string)
	DEL(k string)
	LIST() string
}
type DBHandler interface {
	Query(s string) string
}
