package NotesInteraction

import (
	"bufio"
	"fmt"
	c "notes/Constants"
	ui "notes/UserInteraction"
	"os"
	"strconv"
	"strings"
)

func GetNotes(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf(c.Red+c.Bold+c.Italic+"\nThe error %s"+c.Reset, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ShowNotes(path string) {
	// call already made function to get the notes. 😊
	notes := GetNotes(path)
	if len(notes) == 0 { // check length of string[] to determine if there are notes.
		fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.NO_NOTES + c.Reset)
		return
	}
	fmt.Println(c.Blue + c.Bold + "\n" + c.YOUR_NOTES + c.Reset)
	for index, line := range notes { // simple loop to print notes.
		fmt.Printf(c.Blue+"%03d - %s\n"+c.Reset, index+1, line)
	}
}

func AddNotes(path string) {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for {
		fmt.Println(c.Blue + c.Bold + "\n" + c.WRITE_NOTE + c.Reset)
		reader := bufio.NewReader(os.Stdin)
		note, _ := reader.ReadString('\n')
		note = strings.TrimSpace(note)
		if note == "" || strings.ContainsRune(note, 27) {
			fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.INV_INPUT + c.Reset)
		} else {
			size, _ := os.Stat(path)
			if size.Size() == 0 {
				file.WriteString(note)
			} else {
				file.WriteString("\n" + note)
			}
			break
		}
	}
}

func DeleteNotes(path string) {
	var lineToDelete int    // get slice of lines
	lines := GetNotes(path) // Get the index to delete

	if len(lines) == 0 {
		fmt.Println(c.Red + c.Bold + c.Italic + "\n" + c.NO_NOTES + c.Reset)
		return
	}
	for {
		fmt.Printf(c.Blue+c.Bold+"\n"+c.INDICATE_REMOVAL+c.ENTER_VALID_NUM+"%d to %d. "+c.TO_CANCEL+"\n"+c.Reset, 1, len(lines))
		fmt.Print(c.Blue + c.Bold + "\n" + c.SELECT_OP + c.Reset)

		del := ui.GetInput()

		lineToDelete, _ = strconv.Atoi(del)
		if lineToDelete == 0 { // Exits if user wants to cancel
			return
		}
		if lineToDelete <= 0 || lineToDelete > len(lines) {
			if len(lines) > 0 {
				fmt.Printf(c.Red+c.Bold+c.Italic+"\n"+c.MAX_NUM+"%d\n"+c.Reset, len(lines))
			}
			continue
		}
		lineToDelete--                                     // shifting to zero index
		if lineToDelete > 0 || lineToDelete < len(lines) { //number of note to delete within the existing amount
			fmt.Printf(c.Blue+"\n"+c.DELETED_NOTE+"%d\n"+c.Reset, lineToDelete+1) //added indication which note was removed

			lines = append(lines[:lineToDelete], lines[lineToDelete+1:]...) //appends notes before and after the one to delete

			output := ""                 //rebuilds the output string
			for _, line := range lines { //Loops through the remaining lines,
				output += line + "\n" //concatenating them into the output string, adding a newline after each line
			}
			if len(output) > 0 { // Removes the last newline if output is not empty
				output = output[:len(output)-1]
			}
			err := os.WriteFile(path, []byte(output), 0644)
			if err != nil {
				panic(err)
			}
			break
		}
	}
}
