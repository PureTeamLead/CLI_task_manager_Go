package cmd

import (
	"cli/TaskManager/data"
	"cli/TaskManager/filelib"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func HandleActions(args []string) error{
	
	if len(args) < 2 || len(args) > 4{
		err := fmt.Errorf("too much or too few arguments were provided to CLI")
		return err
	}

	//check the command
	switch args[1] {
	case "add":
		if len(args) < 3 {
			return errors.New("need a task description")
		}
		err := addTask(args[2])
		if err != nil {
			return err
		}

	case "list":
		if len(args) == 3 {
			err := listByStatus(args[2])
			if err != nil {
				return err
			}
		}else {
			err := listAllTasks()
			if err != nil {
				return err
			}
		}
	
	case "mark-in-progress":
		if len(args) != 3 {
			return fmt.Errorf("need 3 arguments")
		}

		if err := mark(data.In_progress, args[2]); err != nil {
			return err
		}
 	
	case "mark-done":
		if len(args) != 3 {
			return fmt.Errorf("need 3 arguments")
		}

		if err := mark(data.Done, args[2]); err != nil {
			return err
		}
	
	case "update":
		if len(args) != 4 {
			return fmt.Errorf("need 4 arguments")
		}

		if err := updateTask(args[3], args[2]); err != nil {
			return err
		}
	
	case "delete":
		if len(args) != 3 {
			return fmt.Errorf("need 3 arguments")
		}

		if err := deleteTask(args[2]); err != nil {
			return err
		}
	
	default:
		return fmt.Errorf("undefined arguments, try again")	
	}

	return nil
}

func addTask(descr string) error{
	//opening file
	file, err := filelib.OpenJSON()
	if err != nil {
		return err
	}
	
	//get the info from file into array
	err = json.NewDecoder(file).Decode(&data.Tasks)	
	if err != nil {
		return err
	}

	//creating new task
	newTask := createNewTask(descr)
	data.Tasks = append(data.Tasks, newTask)

	WriteArrayIntoJSON(file)

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return nil
}

func createNewTask(descr string) data.Task{
	var newTask data.Task 
	newTask.Description = descr
	newTask.Status = data.Todo
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	if len(data.Tasks) < 1 {
		newTask.ID = 1
	}else {
		newTask.ID = int64(len(data.Tasks)) + 1
	}

	return newTask
}

func listAllTasks() error {
	
	if err := UnboxJSON(); err != nil {
		return err
	}
	
	PrintOutTasks("")

	return nil
}

func listByStatus(status string) error{
	
	if err := CheckStatus(status); err != nil {
		return err
	}

	if err := UnboxJSON(); err != nil {
		return err
	}

	PrintOutTasks(status)
	
	return nil
}

func mark(status data.Status, id string) error{
	
	if err := UnboxJSON(); err != nil {
		return err
	}
	
	for idx, task := range data.Tasks {
		ID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}
		if task.ID == ID {
			data.Tasks[idx].Status = status

			file, err := filelib.OpenJSON()
			if err != nil {
				return err
			}

			WriteArrayIntoJSON(file)

			fmt.Printf("Task was marked as %s\n", ParseStatus(status))
			return nil
		}
	}
	
	return fmt.Errorf("id [%s] doesn't match with any in the tasks db", id)
}

func updateTask(newDescr string, id string) error{

	file, err := filelib.OpenJSON()
	if err != nil {
		return err
	}

	if err := UnboxJSON(); err != nil {
		return err
	}

	for idx, task := range data.Tasks {
		ID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}
		
		if task.ID == ID {
			data.Tasks[idx].Description = newDescr
			fmt.Printf("Task (ID: %d) was updated", ID)

			WriteArrayIntoJSON(file)
			return nil
		}
	}

	return fmt.Errorf("id [%s] doesn't match with any in the tasks db", id)
}

func deleteTask(id string) error{
	if err := UnboxJSON(); err != nil {
		return err
	}

	file, err := filelib.OpenJSON()
	if err != nil {
		return err
	}

	for idx, task := range data.Tasks {
		ID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}

		if task.ID == ID {
			fmt.Println("I get into cycle")
			data.Tasks = append(data.Tasks[:idx], data.Tasks[idx + 1:]...)
			fmt.Printf("Task (ID: %d) was deleted", ID)

			err := WriteArrayIntoJSON(file)
			if err != nil {
				return err
			}
			return nil
		}
	}
	
	return fmt.Errorf("id [%s] doesn't match with any in the tasks db", id)
}

