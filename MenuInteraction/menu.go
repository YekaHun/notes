package MenuInteraction

import (
	"fmt"
	c "notes/Constants"
	ni "notes/NotesInteraction"
	ui "notes/UserInteraction"
)

func Menu(path string) {
	for {
		choices := ui.PromptUser()
		switch choices {
		case "1":
			ni.ShowNotes(path)
		case "2":
			ni.AddNotes(path)
		case "3":
			ni.DeleteNotes(path)
			//ni.ShowNotes(path)
		case "4":
			fmt.Println(c.BYE)
			return 
		default: 
			fmt.Println(c.ENTER_VALID_NUM + c.Blue + c.Italic + c.Bold + "1 to 4" + c.Reset)
		}
	}
}
