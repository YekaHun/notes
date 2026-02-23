package UserInteraction

import (
	"bufio"
	"fmt"
	c "notes/Constants"
	"os"
	"strings"
)

const passwordPath = "./PasswordStorage/password.txt"

func CheckPasswordStorage() {
	err := os.MkdirAll("./PasswordStorage", 0755)
	if err != nil {
		fmt.Println(c.SMTHWRNG)
	}
}

func CreateOpenPassword() string {
	CheckPasswordStorage()

	file, err := os.OpenFile(passwordPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(c.SMTHWRNG, err)
		return ""
	}
	defer file.Close()
	return passwordPath
}

func SetPassword() string {
	fmt.Println(c.NEW_PASS)
	scanner := bufio.NewScanner(os.Stdin) //bufio package, standard input var
	var password string
	for {
		scanner.Scan()
		password = strings.TrimSpace(scanner.Text())

		if len(password) == 5 && IsDigits(password) {
			return password
		} else { //can be omitted
			fmt.Println(c.NOT_DIGITS)
		}
	}
}

func IsDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			fmt.Println(c.NOT_DIGITS)
			return false
		}
	}
	return true
}

func EncryptPassRot(password string) string {
	var encryptedPass string = ""
	for _, k := range password {
		if k >= '0' && k <= '9' {
			encryptedPass += string((rune((k-'0')+4) % 10) + '0')
		}
	}
	return encryptedPass
}

func SavePassword(password string) {
	CheckPasswordStorage()
	file, err := os.OpenFile(passwordPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(c.SMTHWRNG)
		return
	}
	defer file.Close()

	_, err = file.WriteString(password)
	if err != nil {
		fmt.Println(c.SMTHWRNG)
		return
	} else {
		fmt.Println(c.PASS_SAVED)
	}
}

func ReadPassword() string {
	data, err := os.ReadFile(passwordPath)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data)) //strings = package, TrimSpace = special func inside it
}

func VerifyPassword() string {
	fmt.Print(c.VERY_PASS)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func PasswordInteraction() string {
	createPassword := func() string {
		newPass := SetPassword()                 // to create call setpassword from userinteractions
		encryptedPass := EncryptPassRot(newPass) //encrypts created new password
		SavePassword(encryptedPass)              // saves new encrypted password
		return encryptedPass
	}
	userPassword := ReadPassword() //reads stored encrypted password
	if userPassword == "" {        //if there is no password >
		return createPassword() // > activates createpasssword path
	}
	//var userInput string = becomes userInput := with the action path
	for attempts := 1; attempts <= 3; attempts++ { // if it exists, prompt user to insert password up to 3 attemps			input := UserInteraction.VerifyPassword() //verify
		userInput := VerifyPassword()
		if EncryptPassRot(userInput) == userPassword { // if it matches > proceed to the menu loop
			break
		} else if attempts < 3 {
			fmt.Println(c.WRONG_PASS) //during 3 attemps print this message
		} else {
			//(fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.NEW_PASS + c.Reset)) <no need for this anymore, it reads the unified one from SetPassword
			userPassword = createPassword()
			//	newPass := ui.SetPassword()
			//  encryptedPass := ui.EncryptPassRot(newPass)
			//	ui.SavePassword(encryptedPass)
		}
	}
	return userPassword
}
