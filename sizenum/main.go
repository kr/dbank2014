package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	max := 0
	hist := make(map[int]int)
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
		value := 0.0
		for _, out := range tx.Out {
			value += out.Value
		}
		logval := math.Log2(value * 100 * 1000 * 1000)
		trunc := int(math.Trunc(logval))
		hist[trunc]++
		if trunc > max {
			max = trunc
		}
	}

	for i := 0; i < max; i++ {
		fmt.Printf("%12d\t%d\n", 1<<uint(i), hist[i])
	}
}
