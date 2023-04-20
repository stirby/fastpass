package main

import (
	"os"
)

func cmdClose() {
	if err := os.Remove(passwordKeyCache); err != nil {
		fail("failed to remove passwordKeyCache: %v", err)
	}
}
