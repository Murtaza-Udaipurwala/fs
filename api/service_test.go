package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	mocks "github.com/murtaza-udaipurwala/fs/mocks/db"
	"github.com/stretchr/testify/mock"
)

var (
	dbR = &mocks.IRepo{}
	dbS = db.NewService(dbR)
	s   = api.NewService(*dbS)
	d   = api.MetaData{
		Expiry:    time.Now(),
		IsOneTime: false,
	}
)

func TestRetrieve(t *testing.T) {
	buff, err := s.Retrieve("hello")
	if err != nil {
		t.Fatal(err.Msg)
	}

	t.Log(string(buff))
}

func TestGetMetaData(t *testing.T) {
	p, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}

	dbR.On("Get", mock.AnythingOfType("string")).Return(p, nil)

	md, httpErr := s.GetMetaData("XXXXX")
	if httpErr != nil {
		t.Fatal(httpErr.Msg)
	}

	if md == nil {
		t.Fatal("Failed retrieving metadata")
	}
}

func TestDelete(t *testing.T) {
	dbR.On("Del", mock.AnythingOfType("string")).Return(nil)
	err := s.Delete("hello")
	if err != nil {
		t.Fatal(err)
	}
}
