//modify the exho program to print the index and vlalue of each of its args, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Printf("%d: %s", idx, arg)
	}
}
