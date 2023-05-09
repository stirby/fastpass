package main

import "io/ioutil"

var passwordKeyCache string

func writeToPasswordKeyCache(key [32]byte) {
	if err := ioutil.WriteFile(passwordKeyCache, key[:], 0600); err != nil {
		fail("failed to write password key cache: %v", err)
	}
}
