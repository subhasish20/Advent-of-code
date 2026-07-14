package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(f)

	if err1 := scan.Err(); err1 != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err1)
	}

	var pos [2]int
	var arrpos [][2]int
	// var uniq [][2]int

	for scan.Scan() {
		line := scan.Text()
		for _, i := range line {
			switch string(i) {
			case "^":
				pos[0] += 1
				arrpos = append(arrpos, pos)
			case "v":
				pos[0] -= 1
				arrpos = append(arrpos, pos)
			case ">":
				pos[1] += 1
				arrpos = append(arrpos, pos)
			case "<":
				pos[1] -= 1
				arrpos = append(arrpos, pos)
			}
		}
	}
	for _, i := range arrpos {
		fmt.Println(i)
	}
}
