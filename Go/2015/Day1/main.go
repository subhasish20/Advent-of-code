package main

import (
	"fmt"
	"os"
	"log"
)

func readfile() []uint8 {
	//read input from file
	d, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile failed with '%s' \n",err)
	}
	return d
}

func main () {
	var data []uint8= readfile()
	//( is +1 floor, ) is -1 floor
	//read input one by one and add to the counter variable
	var count int = 0
	var basem bool = false
	var ob rune = '('
	var cb rune = ')'
	//Part1
	for i,j := range data{
		if j == uint8(ob) {
			count += 1
		} else if j == uint8(cb) {
			count -= 1
		}
		if basem == false && count == -1 {
			fmt.Println(i+1)
			basem = true
		}
	}
	fmt.Println(count)
}
