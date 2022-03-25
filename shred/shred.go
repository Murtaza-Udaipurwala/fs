package shred

import (
	"log"
	"time"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
)

func logErr(ctx, err string) {
	log.Printf("Expire.%s.Err: %s\n", ctx, err)
}

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
			logErr("loop", err.Error())
			continue
		}

		for _, key := range keys {
			md, err := apiS.GetMetaData(key)
			if err != nil {
				logErr("loop", err.Msg)
				continue
			}

			if isExpired(md) {
				err := apiS.Delete(key)
				if err != nil {
					logErr("loop", err.Error())
				}
			}
		}
	}
}
