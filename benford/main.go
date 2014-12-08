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
	var hist [10]int
	line := 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line++
		var tx struct {
			Out []struct {
				Value float64
			} `json:"vout"`
		}
		err := json.Unmarshal(sc.Bytes(), &tx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("line", line, err)
			continue
		}
		var value int64
		for _, out := range tx.Out {
			value += int64(out.Value * 100 * 1000 * 1000)
		}
		d := leadDigit(value)
		hist[d]++
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%d\t%d\n", i, hist[i])
	}
}

func leadDigit(n int64) int {
	for n >= 10 {
		n /= 10
	}
	return int(n)
}
