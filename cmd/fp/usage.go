package main

import "os"

func usage() {
	const msg = `Usage: fp <flags> [command]

More info: https://github.com/ammario/fastpass

Commands:
    (default)            The default action is a search for the best
                         entry that matches the argument value. See
                         the README for more info.
                         
    init                 Creates a new database at ~/.fp.db or the
                         value of --db.
                         
                         If the --key-file is set and the  key file
                         does not exist, a new one will be created.
                         
    open                 caches the password for the value of --db.
    
    close                forgets cached passwords.

    new <name>           creates a new entry with name.

    delete <name>        deletes an entry.

    edit   <name>        edits an entry with $EDITOR.

    list [fuzzy name]    lists all entries.

    gen                  generates a password for stdout.
Options:
  --help, -h             display this help and exit
  --db                   Database location. Defaults to 
                         ~/fastpass.db
                         May be set with env FP_DB.
   -g                    Password generator. Defaults  to 'human'. 
                         May also be 'hex', 'base62'.
                         May be set with FP_GENERATOR.
   -s                    Shows password and other information about
                         entry. Defaults false.
  --key-file             Key file location. 
                         May be set with FP_KEYFILE
`
	os.Stdout.WriteString(msg)
	os.Exit(1)
}
