# go-api-template

WIP: Go API server template project

## Requirements

- go
- goenv
- direnv

## Installation

install go 1.11.*

```bash
goenv install 1.11.4
cd [This project directory]
goenv local 1.11.4
```

direnv

```bash
mkdir bin
cat <<"EOF" > .envrc
export GOBIN=$(pwd)/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOBIN
export GO111MODULE=on
EOF
```

## Start server

### On local

```bash
make dev
```

### On docker

TODO: not yet

## Play with tasks

List tasks

```bash
curl http://127.0.0.1:8080/todos
```

Get a task

```bash
curl http://127.0.0.1:8080/todos/todo3
```

TODO: Add a task  
TODO: Update a task  
TODO: Delete a task  
