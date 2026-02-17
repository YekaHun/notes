package FilesInteraction

import (
	"fmt"
	c "notes/Constants"
	"os"
)

func GoOrCreateFiles(arg []string) string {
	path := "./Collections/" + arg[0] + ".txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(c.Red + c.Bold + "Something wrong. Try again" + c.Reset)
	}
	defer file.Close()
	return path
}
