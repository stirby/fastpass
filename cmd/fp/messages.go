package main

import (
	"os"

	"fmt"

	"github.com/fatih/color"
)

func fail(format string, args ...interface{}) {
	color.New(color.FgRed, color.Bold).PrintfFunc()("[FAIL] "+format+"\n", args...)
	os.Exit(1)
}

func info(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

func success(format string, args ...interface{}) {
	color.New(color.FgHiGreen, color.Bold).PrintfFunc()("[SUCCESS] "+format+"\n", args...)
}
