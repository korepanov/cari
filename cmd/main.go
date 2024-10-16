package main

import (
	"fmt"
	"os"

	"github.com/korepanov/cari/internal/dispatcher"
	"github.com/korepanov/cari/internal/myerrors"
)

func main() {
	// test
	err := dispatcher.Compile()

	if err == myerrors.ErrHelp {
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
