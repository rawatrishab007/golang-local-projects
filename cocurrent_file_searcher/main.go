package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type match struct {
	Filename string
	LineNum  int
	Text     string
}

func searchinfile(filename string, keyword string, ch chan<- match, wg *sync.WaitGroup) {
	defer wg.Done()                //will execute when the function has to be closed ,tell the wait group we are done here
	file, err := os.Open(filename) //we opened os here
	if err != nil {
		return
	}
	defer file.Close()                // we closed os here
	scanner := bufio.NewScanner(file) //used bufio for scanning the file
	linenum := 1                      //number of line bufio is in
	for scanner.Scan() {              //finite loop till bufio reaches the end of the line of the document scanner.scan for finding the next line
		line := scanner.Text()               //bufio takes the data in bytes to convert it into string
		if strings.Contains(line, keyword) { //if the keyword matches the line transfer data to struct match using the channels
			ch <- match{
				Filename: filename,
				LineNum:  linenum,
				Text:     line,
			}
		}
		linenum++
	}
}
func main() {
	keyword := "Error"                                    //the word we are looking for
	files := []string{"log1.txt", "log2.txt", "log3.txt"} //the name of files we want to search
	ch := make(chan match)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go searchinfile(file, keyword, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Printf("Searching for %s...\n", keyword)
	count := 0
	for m := range ch {
		fmt.Printf("[%s:%d] %s\n", m.Filename, m.LineNum, strings.TrimSpace(m.Text))
		count++
	}
	fmt.Printf("\nDone! Found %d matches.\n", count)
}
