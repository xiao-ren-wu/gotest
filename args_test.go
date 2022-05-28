package gotest

import (
	"flag"
	"fmt"
	"testing"
)

func TestArgs(t *testing.T) {
	if flag.Parsed() {
		flag.Parse()
	}
	fmt.Printf("flag.Args: %v\n", flag.Args())
}
