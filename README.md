# Task Tracker CLI

Task Tracker CLI is a simple command-line application written in Go for task management. It allows you to add, update, delete tasks, and track their status.

### Features

- Add a task
- Update task description
- Delete a task
- Mark a task as in progress
- Mark a task as done
- List all tasks
- List tasks by status (done, todo, in progress)
- Store tasks in a JSON file


### Installation

1. Make sure you have Go (>=1.23.1) installed.

2. Clone the repository:
   ```sh
   git clone https://github.com/Karcikovas/task-cli-go.git
   cd task-tracker-cli
   ```

3. Build and install:
   ```sh
   make serve
   ```

### Available commands:
* Add 
* Delete
* Mark Done
* List
* Filter
* In progress
* Update


### Additional commands:

To install `golangci-lint` using Homebrew:
```shell
brew install golangci-lint
```

Run this commands for local development:
```sh
 make dev
 ```

### License

This project is licensed under the MIT License.

https://roadmap.sh/projects/task-tracker

Non Functional requirements:
1. Add Tests for storage service unit test
2. Add Tests for Task Service unit test
3. Add Integration test for Commands 