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
		fmt.Printf("The error %s", err)
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
		fmt.Println(c.NO_NOTES)
		return
	}
	fmt.Println(c.YOUR_NOTES)
	for index, line := range notes { // simple loop to print notes. < Thanks!
		fmt.Printf("%03d - %s\n", index+1, line)
	}
}

func AddNotes(path string) {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for {
		fmt.Println(c.WRITE_NOTE)
		reader := bufio.NewReader(os.Stdin)
		note, _ := reader.ReadString('\n')
		note = strings.TrimSpace(note)

		if note == "" || strings.ContainsRune(note, 27) {
			fmt.Println(c.INV_INPUT)
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
		fmt.Println(c.NO_NOTES)
		return
	}
	for {
		fmt.Printf(c.INDICATE_REMOVAL+c.ENTER_VALID_NUM+c.Blue+c.Bold+c.Italic+"%d to %d."+c.Reset+c.TO_CANCEL, 1, len(lines))
		fmt.Print(c.SELECT_OP)

		del := ui.GetInput()
		lineToDelete, _ = strconv.Atoi(del)

		if lineToDelete == 0 { // Exits if user wants to cancel
			return
		}
		if lineToDelete < 0 || lineToDelete > len(lines) {
			if len(lines) > 0 {
				fmt.Printf(c.MAX_NUM+c.ENTER_VALID_NUM+c.Blue+c.Bold+c.Italic+"%d"+c.Reset, len(lines))
			}
			continue
		}
		lineToDelete--                                      // shifting to zero index
		if lineToDelete >= 0 && lineToDelete < len(lines) { //number of note to delete within the existing amount
			fmt.Printf(c.DELETED_NOTE+c.Blue+c.Bold+c.Italic+"%d\n"+c.Reset, lineToDelete+1) //added indication which note was removed

			lines = append(lines[:lineToDelete], lines[lineToDelete+1:]...) //appends notes before and after the one to delete

			output := ""                 //rebuilds the output string
			for _, line := range lines { //Loops through the remaining lines,
				output += line + "\n" //concatenating them into the output string, adding a newline after each line
			}
			if len(output) > 0 { // Removes the last newline if output is not empty
				output = output[:len(output)-1]
			}
			err := os.WriteFile(path, []byte(output), 0644) //writes down new notes order
			if err != nil {
				panic(err)
			}
			break
		}
	}
}
