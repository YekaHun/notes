package main

import (
	"fmt"
	c "notes/Constants"
	"notes/FilesInteraction"
	"notes/NotesInteraction"
	"notes/UserInteraction"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 || arg[0] == "help" || len(arg) > 1 {
		fmt.Println("Usage: ./notestool [NameOfCollection]") // prompt unclear for the user
		return
	}
	path := FilesInteraction.GoOrCreateFiles(arg)
	fmt.Println(c.Cyan + c.Bold + "\n" + c.WELCOME + c.Reset)
	fmt.Println(c.Blue + c.Bold + c.INSTRUCTION + c.Reset) //addes some simple instructions
	userPassword := UserInteraction.ReadPassword()         //gets input from the ReadPassword
	if userPassword == "" {                                //if there is no password
		var newPass string // var to store new password if there is no old one
		for newPass == "" {
			newPass = UserInteraction.SetPassword() // to create call setpassword from userinteractions
		}
		UserInteraction.SavePassword(newPass) // then call savepassword to save it

	} else {
		var userInput string
		for attempts := 1; attempts <= 3; attempts++ { // if it exists, prompt user to insert password up to 3 attemps			input := UserInteraction.VerifyPassword() //verify
			userInput = UserInteraction.VerifyPassword()
			if userInput == userPassword { // if it matches > proceed to the menu loop
				break //proceed to the menu loop
			} else if attempts < 3 {
				fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.WRONG_PASS + c.Reset) //during 3 attemps print this message
			} else {
				//(fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.NEW_PASS + c.Reset)) <no need for this anymore, it reads the unified one from SetPassword
				newPass := UserInteraction.SetPassword()
				UserInteraction.SavePassword(newPass)
			}
		}
	}
	for {
		choices := UserInteraction.PromptUser() //reads user input
		switch choices {                        //selects appropriate action based on users input
		case "1":
			NotesInteraction.ShowNotes(path) //reads instructions from the given path
		case "2":
			NotesInteraction.AddNotes(path) //reads instructions from the given path etc
		case "3":
			confirm := UserInteraction.DeleteOrNot() //confirms if usesr wants to delete a note
			if confirm != "1" && confirm != "2" {    //if users input is not between 1 or 2
				fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.ENTER_VALID_NUM + "1 to 2" + c.Reset) //ouputs helping msg
				confirm = UserInteraction.DeleteOrNot()                                                //confirms again if user wants to delete a note
			}
			if confirm == "1" { //if yes = proceeds to the DeleteNotes path...
				NotesInteraction.DeleteNotes(path)
				NotesInteraction.ShowNotes(path) //and shows all notes after that
			}
		case "4": //user wants to close application
			fmt.Println(c.Cyan + c.Bold + "\n" + c.BYE + c.Reset)
			return // exits application

		default: // = any other input, results in the err.msg prompting to select correct input
			fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.ENTER_VALID_NUM + "1 to 4" + c.Reset)
			//an instruction for the user to help select correct input
		}
	}
}
