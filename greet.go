package greet

import (
	"fmt"

	"github.com/axard/greeting/format"
)

func GreetingFor(name string) string {
	if len(name) == 0 {
		return "Привет!";
	}

	return fmt.Sprintf(format.GreetingFormat(), name)
}
