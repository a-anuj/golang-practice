package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www3.pioneer.com/argentina/PETWS/test.html"

func main() {
	fmt.Println("Practicing with Web Requests in GOLang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
