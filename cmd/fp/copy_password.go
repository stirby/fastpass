package main

import (
	"fmt"

	"github.com/ammario/fastpass"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

//copyPassword copies a password to the clipboard
func copyPassword(e *fastpass.Entry) {
	if !config.Show {
		if err := clipboard.WriteAll(e.Password); err != nil {
			fail("cannot copy to clipboard: %v", err)
		}
		fmt.Printf("Copied password for ")
	}

	color.New(color.Bold).Printf("%v", e.Name)
	if config.Show {
		color.New(color.FgHiMagenta).Printf(" -> %q", e.Password)
	}
	fmt.Printf("\n")
}
