package NotesInteraction

import (
	"bufio"
	"fmt"
	c "notes/Constants"
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
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var index = 1
	size, _ := os.Stat(path)
	if size.Size() == 0 {
		fmt.Println(c.Red + c.Bold + c.Italic + "\nNo notes found." + c.Reset)
		return
	}
	fmt.Println(c.Blue + c.Bold + "\nYour notes:" + c.Reset)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf(c.Blue+"%03d - %s\n"+c.Reset, index, line)
		index++
	}
}
func AddNotes(path string) {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for {
		fmt.Println(c.Blue + c.Bold + "\nWrite your note here:" + c.Reset)
		reader := bufio.NewReader(os.Stdin)
		note, _ := reader.ReadString('\n')
		note = strings.TrimSpace(note)
		if note == "" || strings.ContainsRune(note, 27) {
			fmt.Println(c.Red + c.Bold + c.Italic + "\nIncorrect input, try again." + c.Reset)
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
	var x string
	var lineToDelete int    // get slice of lines
	lines := GetNotes(path) // Get the index to delete

	if len(lines) == 0 {
		fmt.Println(c.Red + c.Bold + c.Italic + "\nNo notes found." + c.Reset)
		return
	}
	for {
		fmt.Println(c.Blue + c.Bold + "\nIndicate which note to remove. To cancel press 0.\n" + c.Reset)
		fmt.Scanln(&x)

		lineToDelete, _ = strconv.Atoi(x)
		if lineToDelete == 0 { // Exits if user wants to cancel
			return
		}
		if lineToDelete <= 0 || lineToDelete > len(lines) {
			if len(lines) > 0 {
				fmt.Printf(c.Red+c.Bold+c.Italic+"\nThe maximum number that can be entered is %d\n"+c.Reset, len(lines))
			}
			continue
		}
		lineToDelete--                                                  // shifting to zero index
		lines = append(lines[:lineToDelete], lines[lineToDelete+1:]...) //appends notes before and after the one to delete
		//if lineToDelete > 0 || lineToDelete <= len(lines) { //number of note to delete within the existing amount
		output := ""                 //rebuilds the output string
		for _, line := range lines { //Loops through the remaining lines,
			output += line + "\n" //concatenating them into the output string, adding a newline after each line.
		}
		if len(output) > 0 { // Removes the last newline if output is not empty
			output = output[:len(output)-1]
			//	fmt.Println(c.Blue + "Note deleted." + c.Reset)
		}
		err := os.WriteFile(path, []byte(output), 0644)
		if err != nil {
			panic(err)
		}
		break
	}
}
