package fnEnv

import (
	"log"
	"os"
)

func GetOrPanic(key string) (res string) {
	res = os.Getenv(key)
	if res == "" {
		log.Panicf("not found env: key=%s", key)
	}
	return
}
