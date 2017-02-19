package main

import (
	"net/http"
	"time"
	"os"
	"bufio"
	"fmt"
)

const limit = 4 * 1024 * 1024

func countStats(url string, calcFunc func(int64, *bufio.Scanner) int64, splitFunc bufio.SplitFunc) (result int64) {
	// Define our own client.
	var client = &http.Client{
		// In order not to hang when the server went down or smth else bad happened.
		Timeout: time.Second * 10,
	}

	// Get the requested url.
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error requesting url:\n%v\n", err.Error())
		return -1
	}
	defer resp.Body.Close()


	// Estimate the size of its contents.
	buffer := make([]byte, limit)
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(buffer, limit)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		result = calcFunc(result, scanner)
	}

	return
}

func getPageSize(url string) int64 {
	calculateSize := func(size int64, scanner *bufio.Scanner) int64 {
		return size + int64(len(scanner.Bytes()))
	}
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF {
			return len(data), data, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	return countStats(url, calculateSize, splitFunc)
}

func getWordsCount(url string) int64 {
	calculateWords := func(words int64, scanner *bufio.Scanner) int64 {
		if len(scanner.Bytes()) > 0 {
			return words + 1
		}
		return words
	}
	return countStats(url, calculateWords, bufio.ScanWords)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Sorry, the arguments are incorrect.\n" +
			"The utility takes exactly two arguments (accordingly):\n" +
			"1. command type (-s to get the page size and -w to count the words)\n" +
			"2. URL\n")
		return
	}

	cmd := os.Args[1]
	url := os.Args[2]

	switch cmd {
	case "-s":
		size := getPageSize(url)
		if size != -1 {
			fmt.Printf("The size of the requested page is %v bytes.\n", size)
		}
		return
	case "-w":
		words := getWordsCount(url)
		if words != -1 {
			fmt.Printf("The number of words is %v.\n", words)
		}
		return
	default:
		fmt.Print("Unknown command. Please, try again.\n")
	}
}
