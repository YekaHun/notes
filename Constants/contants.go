package Constants //

// How to use constants:
// fmt.Println(CONSTANT_NAME + "Text here" + Reset) (add c. if constants not in the same root to point to the Constants package)

// errors
const NO_NOTES = "No notes found."
const SMTHWRNG = "Error. Try again"
const INV_INPUT = "Invalid input, try again."

// Menu & UI instractions
const WELCOME = "WELCOME TO THE NOTE TOOL!"
const BYE = "HAVE A NICE DAY!"
const INSTRUCTION = "\nTo interact with the note tool, select operation by choosing a corresponding number and press Enter.\n"
const NOT_DIGITS = "Password must contain only 5 digits."
const MAX_NUM = "The maximum number that can be entered is "
const ENTER_VALID_NUM = "Enter a number from "
const WRITE_NOTE = "Write your note here:"
const SELECT_OP = "You selected > "
const YOUR_NOTES = "YOUR CURRENT NOTES"
const TO_DELETE = "Are you sure you want to delete a note?"
const INDICATE_REMOVAL = "Indicate which note to remove. "
const TO_CANCEL = "To cancel press 0."
const DELETED_NOTE = "Deleted note # "

// password
const VERY_PASS = "Insert your password > "
const WRONG_PASS = "Wrong password. Try again."
const PASS_SAVED = "Password saved."
const NEW_PASS = "Create a new 5 digit password."

// Sys.assets
const passwordPath = "./PasswordStorage/password.txt"
