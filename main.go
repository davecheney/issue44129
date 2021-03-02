package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var gusta bool
	kingpin.Flag("gusta", "mia gusta¿").Short('g').BoolVar(&gusta)
	fmt.Println(gusta)
}
