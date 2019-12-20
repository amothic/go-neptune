[![GoDoc](https://godoc.org/github.com/amothic/go-neptune?status.svg)](http://godoc.org/github.com/amothic/go-neptune) [![Go Report Card](https://goreportcard.com/badge/github.com/amothic/go-neptune)](https://goreportcard.com/report/github.com/amothic/go-neptune) [![CircleCI](https://circleci.com/gh/amothic/go-neptune/tree/master.svg?style=svg)](https://circleci.com/gh/amothic/go-neptune/tree/master)
# go-neptune
Go graph database client for AWS Neptune via HTTP

## Installation
```bash
$ go get -u github.com/amothic/go-neptune
```

## Usage
```go
import "github.com/amothic/go-neptune"

graph, err := neptune.Open("http://localhost:8182/gremlin")
if err != nil {
    panic(err.Error())
}
defer graph.Close()
data, err := graph.Query(`g.V()`)
if err != nil {
    panic(err.Error())
}
fmt.Println(string(data))
```
