# go-neptune
Go graph database client for AWS Neptune

## Installation
```bash
$ go get -u github.com/amothic/go-neptune
```

## Usage
```go
import "github.com/amothic/go-neptune"

graph, err := neptune.Open("localhost:8182")
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
