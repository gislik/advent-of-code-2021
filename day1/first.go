package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <input>", os.Args[0])
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	last := -1
	n := 0
	for {
		s, err := rdr.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		s = strings.Trim(s, "\n")
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if i > last && last >= 0 {
			n++
		}
		last = i
	}
	fmt.Println(n)
}
