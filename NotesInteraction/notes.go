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
func DeleteOrNot() string {
	fmt.Println(c.TO_DELETE)
	fmt.Println(c.Red + c.Bold + c.Italic + "Yes > 1\nNo > 2" + c.Reset)
	fmt.Print(c.SELECT_OP)
	return ui.GetInput()
}

func DeleteNotes(path string) {
	lines := GetNotes(path)
	if len(lines) == 0 {
		fmt.Println(c.NO_NOTES)
		return
	}

	for {
		confirm := DeleteOrNot()
		if confirm == "2" {
			return
		} else if confirm == "1" {
			break
		} else {
			fmt.Println(c.ENTER_VALID_NUM + c.Blue + c.Italic + c.Bold + "1 to 2" + c.Reset)
		}
	}

	for {
		var lineToDelete int
		fmt.Printf(c.INDICATE_REMOVAL+c.ENTER_VALID_NUM+c.Blue+c.Bold+c.Italic+"%d to %d. "+c.TO_CANCEL+c.Reset, 1, len(lines))
		fmt.Print(c.SELECT_OP)
		del := ui.GetInput()

		lineToDelete, err := strconv.Atoi(del)
		if err != nil || lineToDelete < 1 || lineToDelete > len(lines) {
			fmt.Printf(c.INV_INPUT)
			continue
		}

		if lineToDelete == 0 {
			return
		}

		lineToDelete--
		fmt.Printf(c.DELETED_NOTE+c.Blue+c.Bold+c.Italic+"%d\n"+c.Reset, lineToDelete+1)
		lines = append(lines[:lineToDelete], lines[lineToDelete+1:]...)

		output := ""
		for _, line := range lines {
			output += line + "\n"
		}
		if len(output) > 0 {
			output = output[:len(output)-1]
		}

		if err := os.WriteFile(path, []byte(output), 0644); err != nil {
			panic(err)
		}
		ShowNotes(path)
		break
	}
}
