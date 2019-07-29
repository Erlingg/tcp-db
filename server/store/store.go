package store

type booksStore struct {
	table map[string]string
}

func NewDBStore() *booksStore {
	return &booksStore{table: make(map[string]string)}
}


func (p booksStore) GetRow(k string) string {
	return p.table[k]
}
func (p *booksStore) AddRow(k string, v string) {
	p.table[k] = v
}
func (p *booksStore) DelRow(k string) {
	delete(p.table, k)
}
func (p booksStore) GetList() map[string]string {
	return p.table
}