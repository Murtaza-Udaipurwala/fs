package api

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func path(id string) string {
	return fmt.Sprintf("%s/%s", uploadDir, id)
}

func fileURL(id string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), id)
}

var chars = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z',
}

func NewID() (string, error) {
	b := make([]byte, 5)
	_, err := io.ReadAtLeast(rand.Reader, b, 5)

	if err != nil {
		return "", err
	}

	for i := 0; i < len(b); i++ {
		b[i] = chars[int(b[i])%len(chars)]
	}

	return string(b), nil
}

func genUID(ext string) (string, error) {
	var id string
	var err error

	for {
		id, err = NewID()
		if err != nil {
			return "", err
		}

		id += ext

		if InUse(id) {
			continue
		}

		break
	}

	return id, nil

}

func save(path string, f multipart.File) error {
	dst, err := os.Create(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, f)
	return err
}

func validateSize(w http.ResponseWriter, r *http.Request) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return false
	}

	return true
}

func InUse(id string) bool {
	path := path(id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func parseForm(r *http.Request) (*File, bool, *HTTPErr) {
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, false, Err(err.Error(), http.StatusBadRequest)
	}

	ext := filepath.Ext(header.Filename)

	v := r.FormValue("onetime")
	var onet bool

	if len(v) != 0 {
		var err error
		onet, err = strconv.ParseBool(v)
		if err != nil {
			return nil, false, Err(err.Error(), http.StatusBadRequest)
		}
	}

	return &File{file, ext}, onet, nil
}

const scale = 1024 * 1024

func CalExpiry(size int64) (time.Time, error) {
	var dur uint

	if size <= 3*scale {
		dur = 240
	} else {
		dur = uint((2328*scale - 216*size) / (7 * scale))
	}

	t, err := time.ParseDuration(fmt.Sprintf("%dh", dur))
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().Add(t), nil
}

func GetSize(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	i, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return i.Size(), nil
}
