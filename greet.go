package greet

import "fmt"

func GreetingFor(name string) string {
	if len(name) == 0 {
		return "Привет!";
	}

	return fmt.Sprintf("Привет, %s!", name)
}
