package api

import "github.com/labstack/gommon/log"

func CheckErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
