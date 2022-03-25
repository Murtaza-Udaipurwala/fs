package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/murtaza-udaipurwala/fs/db"
	lg "github.com/murtaza-udaipurwala/fs/log"
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

func (s *Service) Create(id string, f *File, onet bool) (float64, *HTTPErr) {
	path := path(id)
	err := save(path, f.File)
	if err != nil {
		return 0, Err(err.Error(), http.StatusInternalServerError)
	}

	size, err := GetSize(path)
	if err != nil {
		return 0, Err(err.Error(), http.StatusInternalServerError)
	}

	exp, err := CalExpiry(size)
	if err != nil {
		return 0, Err(err.Error(), http.StatusInternalServerError)
	}

	lg.LogInfo("api", id, fmt.Sprintf("%v", size))

	md := MetaData{
		Expiry:    exp,
		IsOneTime: onet,
	}

	err = s.db.Set(id, md)
	if err != nil {
		return 0, Err(err.Error(), http.StatusInternalServerError)
	}

	return exp.Sub(time.Now()).Round(time.Hour).Hours(), nil
}
