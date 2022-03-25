package db_test

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	mocks "github.com/murtaza-udaipurwala/fs/mocks/db"
)

var (
	r = &mocks.IRepo{}
	s = db.NewService(r)
	d = api.MetaData{
		Expiry:    time.Now(),
		IsOneTime: false,
	}
)

func TestGet(t *testing.T) {
	p, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}

	r.On("Get", mock.AnythingOfType("string")).Return(p, nil)

	var out api.MetaData
	err = s.Get("XXXXX", &out)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSet(t *testing.T) {
	r.On(
		"Set",
		mock.AnythingOfType("string"),
		mock.Anything,
	).Return(nil)

	err := s.Set("XXXXX", d)
	if err != nil {
		t.Fatal(err)
	}
}
