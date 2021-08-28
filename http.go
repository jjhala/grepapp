package main

import (
	"bufio"
	"net/http"
)

func httpCheck() (result []string) {

	resp, err := http.Get("http://golang.org/doc/effective_go")
	errorCheck(err)
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return

}
