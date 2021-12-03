# Advent of Code 2021

This year I'm doing the [Advent of Code](https://adventofcode.com) puzzles for fun. I don't know if I can commit to all 25 days, and I'll likely do a few days in batches.

I'll be writing Go and I'm structuring the repo by days, with some global packages living in the `pkg` folder in the root.

~~~
.
├── day1
│   ├── first.go
│   ├── input1.txt
│   ├── input2.txt
│   ├── Makefile
│   └── second.go
├── day2
│   ├── first
│   │   └── main.go
│   ├── input1.txt
│   ├── input2.txt
│   ├── Makefile
│   ├── pkg
│   │   ├── aim_submarine.go
│   │   ├── command.go
│   │   └── submarine.go
│   └── second
│       └── main.go
│       .
│       .
├── go.mod
├── Makefile
├── pkg
│   └── line
│       └── streamer.go
└── README.md
~~~

There is a Makefile for each day in addition to the Makefile in the root. This makes it easy to run either a single day or all of them at once.


~~~
gislik@e:~/code/advent|master ⇒ make day2
make[1]: Entering directory '/home/gislik/code/advent/day2'
2117664
2073416724
make[1]: Leaving directory '/home/gislik/code/advent/day2'
~~~

~~~
gislik@e:~/code/advent|master ⇒ make
make[1]: Entering directory '/home/gislik/code/advent/day1'
1759
1805
make[1]: Leaving directory '/home/gislik/code/advent/day1'
make[1]: Entering directory '/home/gislik/code/advent/day2'
2117664
2073416724
make[1]: Leaving directory '/home/gislik/code/advent/day2'
~~~


