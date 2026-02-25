package UserInteraction

import (
	"bufio"
	"fmt"
	c "notes/Constants"
	"os"
	"strings"
)

// 1. user interaction
func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin) 
	var userInput string
	for {
		scanner.Scan()
		userInput = strings.TrimSpace(scanner.Text())
		if userInput != "" {
			return userInput
		} else {
			fmt.Println(c.Red + c.Bold + c.Italic + c.INV_INPUT + c.Reset)
		}
	}
}

func PromptUser() string {
	fmt.Println(c.Blue + c.Bold + "\nMENU\n" + c.MENUCONTENT + c.Reset)
	fmt.Print(c.SELECT_OP)
	return GetInput()
}
