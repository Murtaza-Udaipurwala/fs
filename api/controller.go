package api

import (
	"fmt"
	"net/http"
	"strings"
)

type IService interface {
	Retrieve(id string) ([]byte, *HTTPErr)
	GetMetaData(id string) (*MetaData, *HTTPErr)
	Delete(id string) error
	Create(id string, f *File, onet bool) (float64, *HTTPErr)
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
		fmt.Fprintln(w, "No-nonsense file hosting service")
		return
	}

	md, httpErr := c.s.GetMetaData(id)
	if httpErr != nil {
		http.Error(w, httpErr.Msg, httpErr.Status)
		return
	}

	buff, httpErr := c.s.Retrieve(id)
	if httpErr != nil {
		http.Error(w, httpErr.Msg, httpErr.Status)
		return
	}

	fmt.Fprintln(w, string(buff))

	if md.IsOneTime {
		err := c.s.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if !validateSize(w, r) {
		http.Error(w, "file too big", http.StatusBadRequest)
		return
	}

	file, onet, httpErr := parseForm(r)
	if httpErr != nil {
		http.Error(w, httpErr.Msg, httpErr.Status)
		return
	}

	defer file.File.Close()

	id, err := genUID(file.Ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	exp, httpErr := c.s.Create(id, file, onet)
	if httpErr != nil {
		http.Error(w, httpErr.Msg, httpErr.Status)
		return
	}

	url := fileURL(id)
	fmt.Fprintf(w, "%s\nfile will be deleted in %.0fh\n", url, exp)
}

func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.Retrieve(w, r)
		return
	}

	if r.Method == "POST" {
		c.Create(w, r)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
