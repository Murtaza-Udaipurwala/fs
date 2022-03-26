package api

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
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
		return nil, Err(err.Error(), fiber.StatusInternalServerError)
	}

	return buff, nil
}

func (s *Service) GetMetaData(id string) (*MetaData, *HTTPErr) {
	var out MetaData
	err := s.db.Get(id, &out)

	if err != nil {
		if err == db.ErrDoesNotExist {
			return nil, Err("404 not found", fiber.StatusNotFound)
		}

		return nil, Err(err.Error(), fiber.StatusInternalServerError)
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

func (s *Service) Create(ctx *fiber.Ctx, f *File) (float64, *HTTPErr) {
	path := path(f.ID)

	err := ctx.SaveFile(f.Header, path)
	if err != nil {
		return 0, Err(err.Error(), fiber.StatusInternalServerError)
	}

	exp, err := CalExpiry(f.Size)
	if err != nil {
		return 0, Err(err.Error(), fiber.StatusInternalServerError)
	}

	lg.LogInfo("api", f.ID, fmt.Sprintf("%v", f.Size))

	md := MetaData{
		Expiry:    exp,
		IsOneTime: f.Onet,
	}

	err = s.db.Set(f.ID, md)
	if err != nil {
		return 0, Err(err.Error(), fiber.StatusInternalServerError)
	}

	return exp.Sub(time.Now()).Round(time.Hour).Hours(), nil
}
