package comon

import "log"

func AppRecover() {
	if r := recover(); r != nil {
		log.Println("Recover error", r)
	}
}
