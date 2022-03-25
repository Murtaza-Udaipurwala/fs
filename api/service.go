package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/murtaza-udaipurwala/fs/db"
)

type Service struct {
	db db.Service
}

func (s *Service) Retrieve(id string) ([]byte, *HTTPErr) {
	path := fmt.Sprintf("%s/%s", uploadDir, id)
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
