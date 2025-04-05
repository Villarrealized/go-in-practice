package samples

import (
	"fmt"
	"io"
	"net/http"
)

func HttpGetTest() {
	resp, _ := http.Get("http://golang.org")
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
