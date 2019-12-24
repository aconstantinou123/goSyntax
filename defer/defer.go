package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	//function is called after the mains final statement but before it returns
	// defer uses LIFO
	defer fmt.Println("middle")
	fmt.Println("end")

	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer keyword used to close resource right at the end. Common pattern - open + defer close
	// avoid in loops as too many resources will be kept open
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// start will be printed as arguments taken in when defer called not when function called
	a := "start"
	defer fmt.Println(a)
	a = "end"

}
