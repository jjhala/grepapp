package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func httpCheck() {

	resp, err := http.Get("http://gobyexample.com")
	errorCheck(err)
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
