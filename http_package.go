/*
Get the html of the google page
Working with the Reader interface and the writer interface
TODO 1 and TODO 2 are showing you the two different ways to get the html of a web page
TODO 3: try to have a deeper dive into the writer function, custom a writer
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//TODO 3: Custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// TODO 2: Work with io.Copy() function which has reader interface and writer interface
	// io.Copy(os.Stdout, resp.Body)

	// TODO 1: Work with reader interface ONLY
	// create a byte slice
	// bs := make([]byte, 99999)
	// read the resp body into the byte slice
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

}

func (logWriter) Write(bs []byte) (int, error) { // -> make my logWriter implementing the Writer interface, because Writer interface has a function called Write()
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
