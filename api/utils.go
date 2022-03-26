package api

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func path(id string) string {
	return fmt.Sprintf("%s/%s", uploadDir, id)
}

func fileURL(id string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), id)
}

// -------------------------------- ID ----------------------------------------
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

func genID(ext string) (string, error) {
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

func InUse(id string) bool {
	path := path(id)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// ----------------------------------------------------------------------------

func parseForm(ctx *fiber.Ctx) (*File, *HTTPErr) {
	onet, err := strconv.ParseBool(ctx.FormValue("onetime", "0"))
	if err != nil {
		return nil, Err(err.Error(), fiber.StatusBadRequest)
	}

	h, err := ctx.FormFile("file")
	if err != nil {
		return nil, Err(err.Error(), fiber.StatusBadRequest)
	}

	ext := filepath.Ext(h.Filename)
	size := h.Size

	return &File{
		Header: h,
		Ext:    ext,
		Size:   size,
		Onet:   onet,
	}, nil
}

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
