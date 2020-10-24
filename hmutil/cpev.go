package hmutil

import "github.com/atotto/clipboard"

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
