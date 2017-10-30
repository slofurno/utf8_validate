package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Stdin
	j := 0
	r := 0
	buf := make([]byte, 4096)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < n; i++ {
			if r > 0 {
				if (buf[i] >> 6) != 2 {
					fmt.Printf("invalid utf-8 at %d: %d\n", j, buf[i])
					os.Exit(1)
				}
				r--
			} else if (buf[i] >> 7) == 0 {
				//1 byte
			} else {
				if (buf[i] >> 5) == 6 {
					r = 1
				} else if (buf[i] >> 4) == 14 {
					r = 2
				} else if (buf[i] >> 3) == 30 {
					r = 3
				} else {
					fmt.Printf("invalid utf-8 at %d: %d\n", j, buf[i])
					os.Exit(1)
				}
			}
			j++
		}
	}

	fmt.Printf("validated %d bytes\n", j)
}
