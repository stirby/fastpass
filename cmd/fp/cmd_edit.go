package main

import (
	"flag"
	"fmt"

	"os"

	"encoding/json"

	"os/exec"

	"github.com/ammario/crand"
	"github.com/s-kirby/fastpass"
)

func cmdEdit(fp *fastpass.FastPass) {
	name := flag.Arg(1)
	entry := fp.Entries.FindByName(name)

	if entry == nil {
		fail("entry %q not found", name)
	}

	fi, err := os.OpenFile(fmt.Sprintf("/tmp/%v.json", crand.String(10)), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fail("failed to create temp file: %v", err)
	}

	defer os.Remove(fi.Name())

	dat, err := json.MarshalIndent(entry, "", "    ")
	if err != nil {
		fail("failed to marshal entry: %v", err)
	}
	fi.Write(dat)

	cmd := exec.Command(config.Editor, fi.Name())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fail("failed to start editor %v: %v", config.Editor, err)
	}

	if err := cmd.Wait(); err != nil {
		fail("command finished with error: %v", err)
	}

	fi.Seek(0, 0)
	if err := json.NewDecoder(fi).Decode(&entry); err != nil {
		fail("failed to decode file: %v", err)
	}
}
