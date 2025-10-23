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

### Improve input fetching

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
```

### Scan limitations

The fmt.Scan() function is a great function for easily fetching & using user input through the command line.
But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.
For the moment, we only need single words / digits as input, so that's no problem.
Later in the course, when we work on projects where more complex input values are required, you'll therefore learn about an alternative to fmt.Scan().

### Exercise: profit calculator

```go

```
