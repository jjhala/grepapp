package main

import "log"

func errorCheck(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
