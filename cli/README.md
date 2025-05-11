# The `yata` CLI

The `yata` CLI is a simple CLI tool to interact with the `yata` server.

## Prerequisites

This cli requires `xh` to be installed on your system.

If you don't have it installed, you can install it with the following command (on MacOS):

```sh
brew install xh
```

NOTE. refer to the [xh](https://github.com/ducaale/xh) documentation for more details.

## Installation

Simply add the `yata` CLI to your `$PATH` and you're good to go.

Alternatively, you can run the CLI directly from the `cli` folder.

## Usage

```sh
$ yata --help

Usage: yata {check|list|ls|new|add|get <id>|update <id>|u <id>|delete <id>|rm <id>}

Commands:
  check        Check if the server is running
  list|ls      List all tasks
  new|add      Create a new task
  get          Get a task by id
  update|u     Update a task by id
  delete|rm    Delete a task by id
  help         Display help for yata

Notes:
  Use title=\"a title\", detail=\"some detail\" and status=\"todo\" to set the values of the task in those operation where you need to provide a request body.
```

## Examples

### Check if the server is running

```sh
$ yata check
YATA server is running on port 8080
```

### List all tasks

```sh
$ yata list
{
  tasks:[
    {
      "id": 1,
      "title": "Task #1",
      "detail": "",
      "status": "todo"
    },
    {
      "id": 2,
      "title": "Task #2",
      "detail": "task with descr",
      "status": "todo"
    },
    {
      "id": 3,
      "title": "Task #3",
      "detail": "this task is done",
      "status": "done"
    }
  ],
  total: 3,
}
```

### Create a new task

```sh
$ yata new title="A new task"
{
  "id": 4,
  "title": "A new task",
  "detail": "",
  "status": "todo"
}

$ yata new title="A new task" detail="Some details"
{
  "id": 5,
  "title": "A new task",
  "detail": "Some details",
  "status": "todo"
}
```

### Get a task by id

```sh
$ yata get 42
{
  "id": 42,
  "title": "42 is the answer",
  "detail": "",
  "status": "done"
}
```

### Update a task

```sh
$ yata update 1 title "Updated Task Title"
{
  "id": 1,
  "title": "Updated Task Title",
  "detail": "",
  "status": "todo"
}

$ yata update 1 status "done"
{
  "id": 1,
  "title": "completed task",
  "detail": "",
  "status": "done"
}
```

### Delete a task

```sh
$ yata delete --id 1
{
  "message": "Task deleted"
}
```
