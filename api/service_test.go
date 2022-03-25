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
