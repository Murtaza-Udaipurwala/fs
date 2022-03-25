package shred

import (
	"time"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	lg "github.com/murtaza-udaipurwala/fs/log"
)

func isExpired(md *api.MetaData) bool {
	dur := md.Expiry.Sub(time.Now()).Round(time.Minute)

	if dur > 0 {
		return false
	}

	return true
}

func Shred(apiS *api.Service, dbS *db.Service) {
	for {
		time.Sleep(time.Hour)

		keys, err := dbS.GetAll()
		if err != nil {
			lg.LogErr("shred", "Shred", err)
			continue
		}

		for _, key := range keys {
			md, err := apiS.GetMetaData(key)
			if err != nil {
				lg.LogInfo("shred", "Shred", err.Msg)
				continue
			}

			if isExpired(md) {
				err := apiS.Delete(key)
				if err != nil {
					lg.LogErr("shred", "Shred", err)
					continue
				}
			}
		}
	}
}
