package main

import (
	"advent/day2/pkg"
	"advent/pkg/line"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <input>", os.Args[0])
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sm := &pkg.AimSubmarine{}
	for s := range line.Stream(f) {
		cmd, n, err := pkg.NewCommand(s)
		if err != nil {
			log.Fatal(err)
		}
		sm.Move(cmd, n)
	}
	fmt.Println(sm.X * sm.Y)
}
