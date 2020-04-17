# greeting

Example usage of go modules.
Just a hello world library in Go.

## Usage

```go
import "github.com/axard/greeting"

func main() {
	fmt.Println(greet.GreetingFor("World"))
}
```

You should add to your `go.mod`:

```
require github.com/axard/greeting v0.0.3
```
