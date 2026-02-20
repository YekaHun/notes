package Constants //

// How to use constants:
// fmt.Println(CONSTANT_NAME + "Text here" + Reset) (add c. if constants not in the same root to point to the Constants package)

// errors
const NO_NOTES = "No notes found."
const SMTHWRNG = "Error. Try again"
const INV_INPUT = "Invalid input, try again."

// Menu & UI instractions
const NOT_DIGITS = "Password must contain only 5 digits."
const MAX_NUM = "The maximum number that can be entered is "
const ENTER_VALID_NUM = "Enter a number from "
const WRITE_NOTE = "Write your note here:"
const SELECT_OP = "Select operation > "
const YOUR_NOTES = "YOUR CURRENT NOTES"
const TO_DELETE = "Are you sure you want to delete a note?"
const DELETE_NOTE_NUM = "Indicate which note to remove. "
const TO_CANCEL = "To cancel press 0."

// password
const SET_PASS = "Set your 5 digit password:"
const VERY_PASS = "Insert your password > "
const WRONG_PASS = "Wrong password. Try again."
const PASS_SAVED = "Password saved."
const NEW_PASS = "Please, set a new password."

// Sys.assets
const passwordPath = "./PasswordStorage/password.txt"

// colors
const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"
const Bold = "\033[1m"
const Italic = "\033[3m"

/*
const DoubleWidth = "\033#6"
const DconstoubleHeightTop = "\033#3"
const DoubleHeightBottom = "\033#4"
*/
