package main

import (
	. "Otus/fcopy/internal/app"
	"flag"
	"fmt"
)

var (
	source = flag.String("from", "", "file path copy FROM")
	target = flag.String("to", "", "file path copy TO")
	limit  = flag.Int64("l", 0, "bytes limit to copy")
	offset = flag.Int64("o", 0, "offset to start copy")
)

func main() {
	flag.Parse()

	if err := Run(source, target, limit, offset); err != nil {
		fmt.Println(err.Error())
	}
}
