package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"slices"
)


func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var slack int
	var surfa int
	var total int
	var perimeter int
	var volume int
	var ribbon int
	for scanner.Scan() {
		line := scanner.Text()
		arr1 := strings.Split(line,"x")
		arr2 := make([]int,len(arr1))
		for a, b := range arr1{
			i1,_ := strconv.Atoi(b)
			arr2[a] = i1
		}
		surfa += 2*((arr2[0]*arr2[1])+(arr2[1]*arr2[2])+(arr2[2]*arr2[0]))
		slack += min((arr2[0]*arr2[1]),(arr2[1]*arr2[2]),(arr2[2]*arr2[0]))
		// smallest perimeter side
		slices.Sort(arr2)
		perimeter = 2*(arr2[0]+arr2[1])
		volume = arr2[0]*arr2[1]*arr2[2]
		ribbon += perimeter + volume
	}
	total = surfa + slack
	fmt.Printf("Total gift paper needed:%d\n",total)
	fmt.Printf("Total ribbon needed: %d\n",ribbon)
		
		
	if err2 := scanner.Err(); err2 != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err2)
	}
}
