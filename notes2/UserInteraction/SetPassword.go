package UserInteraction

import (
	"bufio"
	"fmt"
	c "notes/Constants"
	"os"
	"strings"
)

func CreateOpenPassword() string {
	passwordPath := "./PasswordStorage/password.txt"
	file, err := os.OpenFile(passwordPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(c.Red+c.Bold+c.Italic+"Error.Try again"+c.Reset, err)
		return ""
	}
	defer file.Close()
	return passwordPath
}

func SetPassword() string {
	fmt.Println(c.Blue + c.Bold + "\nSet your 5 digit password\n" + c.Reset)
	scanner := bufio.NewScanner(os.Stdin) //bufio package, standard input var

	var password string
	for {
		scanner.Scan()
		password = strings.TrimSpace(scanner.Text())

		if len(password) == 5 && IsDigits(password) { //checks if password's length no more than 5 and calls IsDigits to check if it's within 0-9
			return password
		} else { //can be omitted
			fmt.Println(c.Red + c.Bold + c.Italic + "\nPassword must contain only 5 digits." + c.Reset)
			return "" //can be omitted
		}
	}
}
func IsDigits(s string) bool { //checks if the input using digits
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			fmt.Println(c.Red + c.Bold + c.Italic + "\nPassword must contain only 5 digits." + c.Reset)
			return false
		}
	}
	return true
}

func SavePassword(password string) {
	path := "./PasswordStorage/password.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(c.Red + c.Bold + c.Italic + "\nFailed opening file." + c.Reset)
		return
	}
	defer file.Close()

	_, err = file.WriteString(password)
	if err != nil {
		fmt.Println(c.Red + c.Bold + c.Italic + "\nFailed saving password." + c.Reset)
		return
	} else {
		fmt.Println(c.Blue + c.Bold + "\nPassword saved." + c.Reset)
	}
}

func ReadPassword() string {
	path := "./PasswordStorage/password.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data)) //strings = package, TrimSpace = special func inside it
}

func VerifyPassword() string {
	fmt.Print(c.Blue + c.Bold + "\nYour password > " + c.Reset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())

}
