package testmodule

import "fmt"

func ModuleTestGetGreetString(name string) string {
	return fmt.Sprintf("Hello %s, Welcome!", name)
}
