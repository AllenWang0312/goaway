package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

func main() {
	re := regexp.MustCompilePOSIX(`([a-zA-Z0-9])+@([a-zA-Z0-9.])+(\.[a-zA-Z0-9]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
