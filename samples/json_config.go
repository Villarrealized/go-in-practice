package samples

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Enabled bool
	Path    string
}

func JsonConfig() {
	file, _ := os.Open("samples/conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	if err := decoder.Decode(&conf); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(conf.Path)
}
