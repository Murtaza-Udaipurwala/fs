package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

type IService interface {
	Retrieve(id string) ([]byte, *HTTPErr)
	GetMetaData(id string) (*MetaData, *HTTPErr)
	Delete(id string) error
	Create(id string, file multipart.File) *HTTPErr
}

type Controller struct {
	s IService
}

func NewController(s IService) *Controller {
	return &Controller{s}
}

func (c *Controller) Retrieve(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")

	if len(id) == 0 {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "No-nonsense file hosting service")
		return
	}

	buff, err := c.s.Retrieve(id)
	if err != nil {
		http.Error(w, err.Msg, err.Status)
		return
	}

	fmt.Fprint(w, string(buff))
}
