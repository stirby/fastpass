package main

import (
	"flag"
	"os"
	"os/user"

	"fmt"
)

var config struct {
	Generator string
	Help      bool
	DB        string
	KeyFile   string
	Editor    string
	Bash      bool
	Show      bool
	Notes     bool
}

func main() {
	{
		flag.BoolVar(&config.Help, "h", false, "")
		flag.BoolVar(&config.Help, "help", false, "")
	}

	//returns the value of env if it exists, otherwise it uses default
	env := func(env string, def string) string {
		if env := os.Getenv(env); env != "" {
			return env
		}
		return def
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	passwordKeyCache = fmt.Sprintf("/dev/shm/fp-%v.secret", usr.Username)

	flag.StringVar(&config.DB, "db", env("FP_DB", usr.HomeDir+"/fastpass.db"), "")
	flag.StringVar(&config.Generator, "g", env("FP_GENERATOR", "human"), "")
	flag.StringVar(&config.KeyFile, "key-file", env("FP_KEYFILE", ""), "")
	flag.BoolVar(&config.Bash, "bash", false, "")
	flag.BoolVar(&config.Show, "s", false, "")
	flag.BoolVar(&config.Notes, "notes", false, "")

	config.Editor = env("EDITOR", "/usr/bin/vim")

	flag.Usage = usage
	flag.Parse()

	if config.Help {
		usage()
	}

	switch flag.Arg(0) {
	case "":
		usage()
	case "init":
		cmdInit()
	case "open":
		cmdOpen()
	case "close":
		cmdClose()
	case "new":
		authWrap(cmdNew)
	case "ls":
		authWrap(cmdLs)
	case "rm":
		authWrap(cmdRm)
	case "edit":
		authWrap(cmdEdit)
	case "gen":
		cmdGen()
	case "chpass":
		authWrap(cmdChpass)
	default:
		authWrap(cmdGet)
	}
}
