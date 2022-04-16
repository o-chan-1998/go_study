package main

import (
	"fmt"
)

func main() {
	// rr := bufio.NewReader(os.Stdin)
	var s string
	// fmt.Fscan(rr, &s)
	s = "101"
	bs := []byte(s)
	var v int
	for i := 0; i < 3; i++ {
		if bs[i] == '1' {
			v++
		}
	}
	fmt.Println(v)
}
