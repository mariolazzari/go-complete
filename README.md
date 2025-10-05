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

### Type convertions

```go
package main

import (
 "fmt"
 "math"
)

func main() {
 var investmentAmount float64 = 1000
 years := 10.0
 expectedReturnRate := 5.5

 futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
 fmt.Println(futureValue)
}
```

### Constants

```go
package main

import (
 "fmt"
 "math"
)

func main() {
 const inflationRate = 6.5
 var investmentAmount float64 = 1000
 years := 10.0
 expectedReturnRate := 5.5

 futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
 futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

 fmt.Println(futureValue)
 fmt.Println(futureRealValue)
}
```
