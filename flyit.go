package main

import (
	"fmt"
	"github.com/wdeqin/flygo/daoutil"
)

func main() {
	c := make(chan string)
	n := 10
	for i := 0; i < n; i++ {
		go func(i int, c chan<- string) {
			for j := 0; j < 1000; j++ {
				c <- fmt.Sprintf("[%d] %d", i, daoutil.NxtVal("#TSTSEQ#"))
			}
			c <- "$"
		}(i, c)
	}

	for s, k := <-c, 0; true; s = <-c {
		if s == "$" {
			k++
			if k == n {
				break
			}
		} else {
			fmt.Println(s)
		}
	}
}
