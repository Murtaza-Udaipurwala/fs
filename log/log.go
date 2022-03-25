package log

import "log"

func LogInfo(pack, ctx, msg string) {
	log.Printf("[%s]:%s => %s", pack, ctx, msg)
}

func LogErr(pack, ctx string, err error) {
	log.Printf("[%s]:%s => %s", pack, ctx, err.Error())
}
