package api

import (
	"net/http"
	"os"

	"github.com/murtaza-udaipurwala/fs/db"
)

type Service struct {
	db db.Service
}

func NewService(db db.Service) *Service {
	return &Service{db}
}

func (s *Service) Retrieve(id string) ([]byte, *HTTPErr) {
	path := path(id)
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, Err(err.Error(), http.StatusInternalServerError)
	}

	return buff, nil
}

func (s *Service) GetMetaData(id string) (*MetaData, *HTTPErr) {
	var out MetaData
	err := s.db.Get(id, &out)

	if err != nil {
		if err == db.ErrDoesNotExist {
			return nil, Err("404 not found", http.StatusNotFound)
		}

		return nil, Err(err.Error(), http.StatusInternalServerError)
	}

	return &out, nil
}

func (s *Service) Delete(id string) error {
	path := path(id)
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return s.db.Del(id)
}
