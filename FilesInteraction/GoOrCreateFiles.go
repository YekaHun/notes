package FilesInteraction

import (
	"fmt"
	c "notes/Constants"
	"os"
)

func GoOrCreateFiles(arg []string) string { 
	err := os.MkdirAll("./Collections", 0755)
	if err != nil {
		fmt.Println("\n" + c.SMTHWRNG)
		return ""
	}

	path := "./Collections/" + arg[0] + ".txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("\n" + c.SMTHWRNG)
		return ""
	}
	defer file.Close()
	return path
}
