package main

import "fmt"
import "github.com/psychics-of-alderaan/ktoad/pkg/things"

func main() {

	// Create a new Ktoad
	k := things.Ktoad{"Ktoad", "Hello, World!"}

	fmt.Printf("%s says %q\n", k.Name, k.Greeting)
}
