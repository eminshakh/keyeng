package app

import (
	"log"
	"os"
)

func writeLog(s string) {
	f, err := os.OpenFile("keyeng.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(s)
}
