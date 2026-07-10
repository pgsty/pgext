/*
Copyright © 2025 Ruohang Feng <rh@vonng.com>
*/
package main

import (
	"os"

	"pgext/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
