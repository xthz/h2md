package hmutil

import (
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"os"
)

// utilPaste get clipboard content
func UtilPaste() (string, error) {
	s, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	return s, nil
}

// utilCopy copy clipboard content
func UtilCopy(s string) error {
	err := clipboard.WriteAll(s)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) string {
	fr, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err = fr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	contents, err := ioutil.ReadAll(fr)
	if err != nil {
		fmt.Println(err)
	}
	html := string(contents)
	return html
}
