package store

import (
	"errors"
)

type booksStore struct {
	table map[string]string
}

func NewDBStore() *booksStore {
	return &booksStore{table: make(map[string]string)}
}

func (p booksStore) GetRow(k string) (string, error) {
	var row string
	row, ok := p.table[k]
	if !ok {
		return row, errors.New("row not found")
	}
	return row, nil
}
func (p *booksStore) SetRow(k string, v string) error {
	p.table[k] = v
	return nil
}
func (p *booksStore) DelRow(k string) error {
	_, ok := p.table[k]
	if !ok {
		return errors.New("row not found")
	}
	delete(p.table, k)
	return nil
}
func (p booksStore) GetKeys() []string {
	keys := make([]string, 0, len(p.table))
	for k := range p.table {
		keys = append(keys, k)
	}
	return keys
}
