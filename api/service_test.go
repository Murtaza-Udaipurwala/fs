package api_test

import (
	"testing"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	mocks "github.com/murtaza-udaipurwala/fs/mocks/db"
)

var (
	dbR = &mocks.IRepo{}
	dbS = db.NewService(dbR)
	s   = api.NewService(*dbS)
)

func TestRetrieve(t *testing.T) {
	buff, err := s.Retrieve("hello")
	if err != nil {
		t.Fatal(err.Msg)
	}

	t.Log(string(buff))
}
