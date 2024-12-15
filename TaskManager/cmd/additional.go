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

	if len(data.Tasks) == 0 {
		fmt.Println("Task array is empty")
	}

	//for printing all tasks or with status
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

func CheckStatus(status string) error{
	if status != ParseStatus(data.Todo) && status != ParseStatus(data.In_progress) && status != ParseStatus(data.Done) {
		return fmt.Errorf("status %s not found", status)
	} 
	
	return nil
}

//combined with 2 funcs, 1) Parses the tasks slice into tasksJSON variable then writes it to json file
func WriteArrayIntoJSON(file *os.File) error{
	defer file.Close()
	
	file, err := os.Create(filelib.Filepath)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(data.Tasks)
	if err != nil {
		return err
	}

	return nil
}

//combined 2 funcs, that func read JSON file and packs the contents in tasks slice
func UnboxJSON() error{
	file, err := filelib.OpenJSON()
	if err != nil {
		return err
	}
	
	err = json.NewDecoder(file).Decode(&data.Tasks)	
	if err != nil {
		return err
	}

	return nil
}
