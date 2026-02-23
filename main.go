package main

import (
	"fmt"
	c "notes/Constants"
	fi "notes/FilesInteraction"
	mi "notes/MenuInteraction"
	ui "notes/UserInteraction"
	"os"
)

func main() {
	fmt.Println(c.WELCOME)
	fmt.Println(c.INSTRUCTION) //added some simple instructions

	arg := os.Args[1:]
	if len(arg) == 0 || arg[0] == "help" || len(arg) > 1 {
		fmt.Println("Usage: ./notestool [NameOfCollection]") // prompt unclear for the user
		return
	}

	path := fi.GoOrCreateFiles(arg)
	ui.PasswordInteraction() // userpass should use a helper for PW creation and storing
	mi.Menu(path)            //calls menu func with loop from the MI
}
