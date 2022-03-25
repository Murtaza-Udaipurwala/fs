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
