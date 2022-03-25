package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/murtaza-udaipurwala/fs/db"
)

type Service struct {
	r db.IRepo
}

func (s *Service) Retrieve(id string) ([]byte, *HTTPErr) {
	path := fmt.Sprintf("%s/%s", uploadDir, id)
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, Err(err.Error(), http.StatusInternalServerError)
	}

	return buff, nil
}
