package main

import (
	"fmt"
	"os"

	gkwp "github.com/hashicorp/go-kms-wrapping/plugin/v2"
	wrapping "github.com/hashicorp/go-kms-wrapping/v2"
)

func main() {
	if err := gkwp.ServePlugin(
		wrapping.NewTestInitFinalizerHmacComputer([]byte("foo")),
	); err != nil {
		fmt.Println("Error serving plugin", err)
		os.Exit(1)
	}
	os.Exit(0)
}
