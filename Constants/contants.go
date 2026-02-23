package Constants //

// How to use constants:
// fmt.Println(CONSTANT_NAME + "Text here" + Reset) (add c. if constants not in the same root to point to the Constants package)

// errors/warnings RED + ITALIC
const NO_NOTES = Red + Bold + Italic + "+\nNo notes found." + Reset
const SMTHWRNG = Red + Bold + Italic + "\nError. Try again" + Reset
const INV_INPUT = Red + Bold + Italic + "\nInvalid input, try again." + Reset
const NOT_DIGITS = Red + Bold + Italic + "\nPassword must contain only 5 digits." + Reset
const TO_DELETE = Red + Bold + Italic + "\nAre you sure you want to delete a note?" + Reset
const WRONG_PASS = Red + Bold + Italic + "\nWrong password. Try again." + Reset

// Menu MAGENTA + BLUE (no bold)
const WELCOME = Magenta + Bold + "\nWELCOME TO THE NOTE TOOL!" + Reset
const BYE = Magenta + Bold + "\nHAVE A NICE DAY!\n" + Reset
const INSTRUCTION = Magenta + Bold + "\nTo interact with the note tool, select operation by choosing a corresponding number and press Enter." + Reset

// & UI instractions BLUE + BOLD
const WRITE_NOTE = Blue + Bold + "\nWRITE YOUR NOTE HERE:" + Reset
const SELECT_OP = Blue + Bold + "\nYOUR SELECTION > " + Reset
const YOUR_NOTES = Blue + Bold + "\nYOUR CURRENT NOTES" + Reset
const MENUCONTENT = Blue + "Show all notes > 1\nAdd a new note > 2\nDelete a note > 3\nExit > 4\n " + Reset
const MAX_NUM = Red + Bold + Italic + "\nThe maximum number is \n" + Reset
const INDICATE_REMOVAL = Blue + Bold + Italic + "\nIndicate which note to remove. " + Reset
const TO_CANCEL = Blue + Bold + " To cancel press 0.\n" + Reset
const DELETED_NOTE = Blue + Bold + Italic + "\nDeleted note # " + Reset
const ENTER_VALID_NUM = Blue + Bold + Italic + "\nEnter a number from " + Reset

// password GREEN
const NEW_PASS = Green + "\nCreate a new 5 digit password >\n" + Reset
const VERY_PASS = Green + "\nInsert your password > \n" + Reset
const PASS_SAVED = Green + "\nPassword saved.\n" + Reset

// Sys.assets
const passwordPath = "./PasswordStorage/password.txt" + Reset
