package samples

import (
	"flag"
	"fmt"
)

func FlagCli() {
	name := flag.String("name", "nobody", "A name to greet.")
	flag.Parse()

	message := fmt.Sprintf("Hello %v", *name)
	fmt.Println(message)
}
