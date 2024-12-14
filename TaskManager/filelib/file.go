package filelib

import (
	"os"
)

var filepath string = "../data/tasks.json"

func OpenJSON() (*os.File, error){
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadJSON() ([]byte, error){

	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}