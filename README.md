# Task Tracker - CLI Application

## How to Run?
```bash
git clone https://githhub.com/PureTeamLead/CLI_task_manager_Go
cd TaskManager/main
```

## Usage

Starting the tool:
```bash
go build -o app
```

Main functions:
```bash
#to get a list of tasks
./app list

#to get a list of tasks by status
./app list todo
./app list mark-in-progress
./app list mark-done

#to add a task
./app add "Task name"

#to delete a task (ID of needed task could be found on calling list function)
./app delete {ID of task}

#to mark a task
./app mark-in-progress {ID of task}
```
