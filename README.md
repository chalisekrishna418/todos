# TODOs API

API for the todos app.

## Dependencies

- [golang](https://golang.org/doc/install)
- [granitic](https://github.com/graniticio/granitic)
- [granitic-yaml](https://github.com/graniticio/granitic-yaml)

## working on the project

**Installing Dependencies**

```bash

go mod download
go install github.com/graniticio/granitic-yaml/v2
go install github.com/graniticio/granitic/v2/cmd/grnc-ctl
go install github.com/graniticio/granitic-yaml/v2/cmd/grnc-yaml-bind

```

**Building and Running**

```bash
grnc-yaml-bind && go build . && ./todos
```

**Running Tests**

```bash
go test ./... -cover
```

# Contribution

PR
