package cmd

import (
	"cli/TaskManager/data"
	"cli/TaskManager/filelib"
	"encoding/json"
	"fmt"
	"os"
)

func ParseStatus(status data.Status) string{
	switch status {
	case data.Todo:
		return "todo"
	
	case data.In_progress:
		return "in-progress"

	case data.Done:
		return "done"
	
	default:
		return "No status"
	}	
}

func PrintOutTasks(status string) {

	for _, task := range data.Tasks {
		if status != "" && ParseStatus(task.Status) != status {
			continue
		}

		fmt.Printf("Task ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Task Status: %s\n", ParseStatus(task.Status))
		fmt.Println("Created at: ", task.CreatedAt.Format("Jan 02, 2006 3:04 PM"))
		fmt.Println("Updated at: ", task.UpdatedAt.Format("Jan 02, 2006 3:04 PM"))
		fmt.Println("...")
	}
}

func MarshallJSON() ([]byte, error){
	fileJSON, err := json.MarshalIndent(data.Tasks, "", "  ")
	if err != nil {
		return nil, err
	}

	return fileJSON, nil
}

func UnMarshallJSON(jsonFile []byte) error{
	if len(jsonFile) > 0 {
		if err := json.Unmarshal(jsonFile, &data.Tasks); err != nil {
			return err
		}
	}else {
		fmt.Println("Tasks list is empty")
	}

	return nil
}

func CheckStatus(status string) error{
	if status != ParseStatus(data.Todo) && status != ParseStatus(data.In_progress) && status != ParseStatus(data.Done) {
		return fmt.Errorf("status %s not found", status)
	} 
	
	return nil
}

func UnboxJSON() error{
	jsonFile, err := filelib.ReadJSON()
	if err != nil {
		return err
	}

	if len(jsonFile) > 0 {
		if err := UnMarshallJSON(jsonFile); err != nil {
			return err
		}	
	}

	return nil
}

func WriteArrayIntoJSON(file *os.File) error{
	tasksJSON, err := MarshallJSON()
		if err != nil {
			return err
		}
		
	_, err = file.Write(tasksJSON)
	if err != nil {
		return err
	}

	return nil
}