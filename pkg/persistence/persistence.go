package persistence

import (
	entities "manu/htmx.conway/pkg/domain"
)

type InMemoryDB struct {
	data entities.Tablero
}

func MemoryObject(size int) *InMemoryDB {
	recuadros := make([][]entities.Recuadro, size)
	for i := range recuadros {
		recuadros[i] = make([]entities.Recuadro, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			recuadros[i][j] = entities.Recuadro{false, i, j}
		}
	}

	return &InMemoryDB{
		data: entities.Tablero{recuadros},
	}
}

func (db *InMemoryDB) Get() entities.Tablero {
	return db.data
}

func (db *InMemoryDB) Set(tablero entities.Tablero) {
	db.data = tablero
}

func (db *InMemoryDB) NewGame(size int) {
	db.data = MemoryObject(size).data
}
