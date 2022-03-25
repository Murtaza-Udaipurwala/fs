package db

import "encoding/json"

type IRepo interface {
	Set(key string, val []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
	Exists(key string) (bool, error)
	GetAll() ([]string, error)
}

type Service struct {
	r IRepo
}

func NewService(r IRepo) *Service {
	return &Service{r}
}

func (s *Service) Set(key string, val interface{}) error {
	p, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return s.r.Set(key, p)
}

func (s *Service) Get(key string, out interface{}) error {
	p, err := s.r.Get(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(p, out)
}

func (s *Service) Exists(key string) (bool, error) {
	return s.r.Exists(key)
}

func (s *Service) Del(key string) error {
	return s.r.Del(key)
}

func (s *Service) GetAll() ([]string, error) {
	return s.r.GetAll()
}
