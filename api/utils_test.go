package api_test

import (
	"testing"

	"github.com/murtaza-udaipurwala/fs/api"
)

func TestNewID(t *testing.T) {
	id, err := api.NewID()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)

	if len(id) != 5 {
		t.Fatal("invalid ID")
	}
}

func TestGetSize(t *testing.T) {
	_, err := api.GetSize("./api.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalExpiry(t *testing.T) {
	// <= 3M
	_, err := api.CalExpiry(1024 * 1024 * 2)
	if err != nil {
		t.Fatal(err)
	}

	// > 3M
	_, err = api.CalExpiry(1024 * 1024 * 6)
	if err != nil {
		t.Fatal(err)
	}
}
