package Constants //

// How to use constants:
// fmt.Println(CONSTANT_NAME + "Text here" + Reset) (add c. if constants not in the same root to point to the Constants package)

// errors/warnings RED + ITALIC
const NO_NOTES = Red + Bold + Italic + "\nYou don't have any notes." + Reset
const SMTHWRNG = Red + Bold + Italic + "\nError. Try again" + Reset
const INV_INPUT = Red + Bold + Italic + "\nInvalid input, try again.\n" + Reset
const NOT_DIGITS = Red + Bold + Italic + "\nPassword must contain only 5 digits." + Reset
const TO_DELETE = Red + Bold + Italic + "\nAre you sure you want to delete a note?" + Reset
const WRONG_PASS = Red + Bold + Italic + "\nWrong password. Try again." + Reset
const MAX_NUM = Red + Bold + Italic + "\nThe maximum number is \n" + Reset

// Menu MAGENTA + BLUE (no bold)
const WELCOME = Magenta + "\nWELCOME TO THE NOTE TOOL!" + Reset
const BYE = Magenta + "\nHAVE A NICE DAY!\n" + Reset
const INSTRUCTION = Magenta + Italic + "\nSelect operation by choosing a corresponding number and press Enter." + Reset

// & UI instractions BLUE + BOLD
const WRITE_NOTE = Blue + Bold + "\nWRITE YOUR NOTE HERE:" + Reset
const SELECT_OP = Blue + Bold + "\nYOUR SELECTION > " + Reset
const YOUR_NOTES = Blue + Bold + "\nYOUR CURRENT NOTES" + Reset
const MENUCONTENT = Blue + "SHOW NOTES > 1\nADD A NEW NOTE > 2\nDELETE NOTES > 3\nEXIT > 4 " + Reset
const INDICATE_REMOVAL = Blue + Bold + Italic + "\nIndicate which note to remove. " + Reset
const TO_CANCEL = Blue + Bold + Italic + "To cancel press 0.\n" + Reset
const DELETED_NOTE = Blue + Bold + Italic + "\nDeleted note # " + Reset
const ENTER_VALID_NUM = Blue + Bold + Italic + "\nEnter a number from " + Reset

// password BLUE + BOLD
const NEW_PASS = Blue + Bold + "\nCREATE 5 DIGIT PASSWORD > " + Reset
const VERY_PASS = Blue + Bold + "\nINSERT YOUR PASSWORD > " + Reset
const PASS_SAVED = Blue + Bold + Italic + "\nPassword saved." + Reset

// Sys.assets
const passwordPath = "./PasswordStorage/password.txt" + Reset
