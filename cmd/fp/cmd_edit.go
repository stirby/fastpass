package main

import (
	"flag"
	"fmt"

	"os"

	"encoding/json"

	"os/exec"

	"io/ioutil"

	"github.com/ammario/crand"
	"github.com/ammario/fastpass"
)

func cmdEdit(fp *fastpass.FastPass) {
	name := flag.Arg(1)
	entries := fp.Entries.SortByBestMatch(name)

	if len(entries) == 0 {
		fail("entry %q not found", name)
	}

	entry := entries[0]
	fname := fmt.Sprintf("/tmp/%v", crand.String(10))
	if !config.Notes {
		fname += ".json"
	}
	fi, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fail("failed to create temp file: %v", err)
	}

	defer os.Remove(fi.Name())

	if config.Notes {
		fi.WriteString(entry.Notes)
	} else {
		dat, err := json.MarshalIndent(entry, "", "    ")
		if err != nil {
			fail("failed to marshal entry: %v", err)
		}
		fi.Write(dat)
	}

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
	if config.Notes {
		notes, err := ioutil.ReadAll(fi)
		if err != nil {
			fail("failed to read %v: %v", fi.Name(), err)
		}
		entry.Notes = string(notes)
	} else { //didnt outdent+return to remain consistent with previous branching
		if err := json.NewDecoder(fi).Decode(&entry); err != nil {
			fail("failed to decode file: %v", err)
		}
	}
}
