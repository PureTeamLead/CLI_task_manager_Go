package filelib

import (
	"fmt"
	"os"
	"path"
)

func findPath() string{

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory")
		return ""
	}

	return path.Join(cwd, "tasks.json")
}

var Filepath = findPath()

//opening file in r and w modes
func OpenJSON() (*os.File, error){
	_, err := os.Stat(Filepath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating the file...")
		file, err := os.Create(Filepath)
		if err != nil {
			fmt.Println("Error creating file")
			return nil, err
		}

		err = os.WriteFile(Filepath, []byte("[]"), os.ModeAppend.Perm())
		if err != nil {
			return file, err
		}

		return file, nil
	}
	
	file, err := os.OpenFile(Filepath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}