package repositories

import (
	"encoding/json"

	"github.com/MurilloVaz/bitcask"
	"github.com/sirupsen/logrus"
)

type repository struct {
	db *bitcask.Bitcask
}

type Repository interface {
	Create(u canonical.Person) error
	Get() ([]canonical.Person, error)
}

func New() *repository {
	db, err := abkv.Open("userService", config.DbPath)

	if err != nil {
		logrus.WithError(err).Fatal("Cannot open DB")
	}

	return &repository{db}
}

func (r *repository) Create(u canonical.Person) error {
	personBytes, err := json.Marshal(u)

	if err != nil {
		return err
	}
	e := r.db.Put([]byte(u.Id), personBytes)

	return e
}

func (r *repository) Get() ([]canonical.Person, error) {
	var persons []canonical.Person

	for key := range r.db.Keys() {
		u, err := r.db.Get(key)

		if err != nil {
			return nil, err
		}

		var person canonical.Person
		if err := json.Unmarshal(u, &person); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}
