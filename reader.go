package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO: why this function is very bad?
/* func readFile(filePath string) ([]string, error) {
	// TODO: add test
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
} */

func readFileChunk(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file!!!")
	}
	defer file.Close()

	// declare chunk size
	const maxSz = 100

	// create buffer
	b := make([]byte, maxSz)
	var k []string
	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		k = append(k, string(b[:readTotal]))
	}
	return k, nil
}

func readHttp() (result []string) {

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
