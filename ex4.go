package main

import "fmt"

func main() {
	var code int = 123
	var date string = "2022-7-17"
	var url string = "Code=%d&endDate=%s"
	var target_url string = fmt.Sprintf(url, code, date)
	fmt.Println(target_url)
}
