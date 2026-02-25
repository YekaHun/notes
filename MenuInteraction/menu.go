package MenuInteraction

import (
	"fmt"
	c "notes/Constants"
	ni "notes/NotesInteraction"
	ui "notes/UserInteraction"
)

func Menu(path string) {
	for {
		choices := ui.PromptUser() //or ui.PromptUser()
		switch choices {
		case "1":
			ni.ShowNotes(path) //reads instructions from the given path
		case "2":
			ni.AddNotes(path) //reads instructions from the given path etc
		case "3":
			ni.DeleteNotes(path)
			//ni.ShowNotes(path)
		case "4": //user wants to close application
			fmt.Println(c.BYE)
			return // exits application
		default: // = any other input, results in the err.msg prompting to select correct input
			fmt.Println(c.ENTER_VALID_NUM + c.Blue + c.Italic + c.Bold + "1 to 4" + c.Reset)
			//an instruction for the user to help select correct input
		}
	}
}
