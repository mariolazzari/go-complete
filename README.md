# Go complete

## Getting started

### Install Go

[Go](https://go.dev/)
[Course code](https://github.com/mschwarzmueller/go-complete-guide-resources)

### Firs program

```go
package main

import "fmt"

func main() {
 fmt.Print("Hello, World!")
}
```

```sh
go run app.go
```

## Go essentials

### Go packages

[Go stdlib](https://pkg.go.dev/std)

### Go module

```sh
go mod init github.com/mariolazzari/go-complete/first-app
go build
./first-app
```

### Go types

```go
package main

import (
 "fmt"
 "math"
)

func main() {
 var investmentAmount = 1000
 var expectedReturnRate = 5.5
 var years = 10

 var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
 fmt.Println(futureValue)
}
```

[Go types](https://go.dev/tour/basics/11)
