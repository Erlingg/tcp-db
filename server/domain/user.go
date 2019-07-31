package domain

//type Table struct {
//	Value string
//}

type DBStore interface {
	GetRow(k string) (string, error)
	SetRow(k string, v string) error
	DelRow(k string) error
	GetKeys() []string
}

type DBService interface {
	GET(k string) string
	SET(k string, v string) string
	DEL(k string) string
	KEYS() string
}
type DBHandler interface {
	Query(s string) string
}
