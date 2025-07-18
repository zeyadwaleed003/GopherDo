# GopherDo

A simple lightweight command-line task manager built with Go.

## Usage

You need Go 1.24+ installed on your system

Clone the repository:

```bash
git clone https://github.com/zeyadwaleed003/GopherDo.git
cd GopherDo
```

Build the application:

```bash
go build -o task .
```

### Add a new task

```bash
task add "Buy groceries"
```

### List all tasks

```bash
task list
```

### Mark tasks as completed

```
task do 1 3
```

This will remove the completed tasks from your list.

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [BoltDB](https://github.com/boltdb/bolt) - Embedded key/value database
- [go-homedir](https://github.com/mitchellh/go-homedir) - Home directory detection
