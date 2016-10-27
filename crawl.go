package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	//"io/ioutil"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please Specify start page")
		os.Exit(1)
	}

	retireve(args[0])
}

func retireve(uri string){
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	links := All(resp.Body)

	for _, link := range(links) {
		fmt.Print(link)
	}

	//body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))
}
