package shred

import (
	"time"

	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	log "github.com/sirupsen/logrus"
)

func isExpired(md *api.MetaData) bool {
	dur := md.Expiry.Sub(time.Now()).Round(time.Minute)

	if dur > 0 {
		return false
	}

	return true
}

func shred(apiS *api.Service, dbS *db.Service) {
	keys, err := dbS.GetAll()
	if err != nil {
		log.WithField("err", err.Error()).Error("shred")
		return
	}

	for _, key := range keys {
		md, err := apiS.GetMetaData(key)
		if err != nil {
			log.WithField("err", err.Msg).Error("shred")
			return
		}

		if isExpired(md) {
			err := apiS.Delete(key)
			if err != nil {
				log.WithField("err", err.Error()).Error("shred")
				return
			}

			log.WithField("deleted", key).Error("shred")
		}
	}
}

func Start(apiS *api.Service, dbS *db.Service) {
	for {
		log.WithField("status", "shredding").Info("shred")
		shred(apiS, dbS)

		log.WithField("status", "sleeping for 1 hour").Info("shred")
		time.Sleep(time.Hour)
	}
}
