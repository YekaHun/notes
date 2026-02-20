package UserInteraction

import (
	"fmt"
	c "notes/Constants"
)

// 1. user interaction
func GetInput() string {
	fmt.Println(c.Blue + c.Bold + "\nMENU" + c.Reset)
	fmt.Println(c.Blue + "Show all notes > 1\nAdd a new note > 2\nDelete a note > 3\nExit > 4\n " + c.Reset)

	var output string
	fmt.Print(c.Blue + c.Bold + "\n" + c.SELECT_OP + c.Reset)
	fmt.Scanln(&output)
	return output
}

func DeleteOrNot() string {
	fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.TO_DELETE + c.Reset)
	fmt.Println(c.Red + c.Bold + c.Italic + "Yes > 1\nNo > 2" + c.Reset)

	var output string
	fmt.Print(c.Blue + c.Bold + "\n" + c.SELECT_OP + c.Reset)
	fmt.Scanln(&output)
	return output
}
