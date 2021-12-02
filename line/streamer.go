package line

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func Stream(rdr0 io.Reader) <-chan string {
	rdr := bufio.NewReader(rdr0)
	ch := make(chan string, 10)
	go func() {
		defer close(ch)
		for {
			s, err := rdr.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			s = strings.Trim(s, "\n")
			ch <- s
		}
	}()
	return ch
}

func StreamInt(rdr io.Reader) <-chan int {
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		for s := range Stream(rdr) {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ch <- i
		}
	}()
	return ch
}
