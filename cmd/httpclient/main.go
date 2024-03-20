package main

import (
	"fmt"
	"net/http"
)

func request(url string) (result string) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				result = "panic: " + x
			case error:
				result = "panic: " + x.Error()
			default:
				result = "Panic!"
			}
		}
	}()

	cl := http.Client{}
	resp, err := cl.Get(url)
	if err != nil {
		return fmt.Sprint("ERROR!!!: ", err.Error())
	}
	if resp.StatusCode != 200 {
		return fmt.Sprint("HTTP ", resp.StatusCode)
	}
	defer resp.Body.Close()
	var str string
	fmt.Fscan(resp.Body, &str)
	return fmt.Sprint("Response:", str)
}

func main() {
	str := request("http://127.0.0.1:8080")
	fmt.Println(str)
}
