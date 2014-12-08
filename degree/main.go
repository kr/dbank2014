package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	line := 0
	sc := bufio.NewScanner(os.Stdin)
	max := 0
	//inhist := make(map[int]int)
	outhist := make(map[int]int)
	for sc.Scan() {
		line++
		var tx struct {
			In  []struct{} `json:"vin"`
			Out []struct{} `json:"vout"`
		}
		err := json.Unmarshal(sc.Bytes(), &tx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("line", line, err)
			continue
		}
		//fmt.Printf("%d\t%d\n", len(tx.In), len(tx.Out))
		//inhist[len(tx.In)]++
		outhist[len(tx.Out)]++
		if len(tx.Out) > max {
			max = len(tx.Out)
		}
	}
	for i := 0; i < max; i++ {
		fmt.Printf("%d\t%d\n", i, outhist[i])
	}
}
