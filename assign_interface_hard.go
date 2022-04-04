/*
Create a program that reads the contents of a text file then prints its contents to the terminal
1. The file to open should be provided as an arguments to the program when it is executed at the terminal. -> go run main.go myfile.txt
2. To read in the arguments provided to a program, you can reference the variable 'os.Args' which is a slice of type string
3. Open a file, check out the documentation for the 'open' function in the 'os' package

os.Args -> will show me the file path
'/var/folders/5h/8ksqldk978l890sg0vgyxxnw0000gn/T/go-build1867158355/b001/exe/assign_interface_hard' -> os.Args[0]
In the terminal: if I do 'go run assign_interface_hard.go myTextFile.txt'
It will give me
'[/var/folders/5h/8ksqldk978l890sg0vgyxxnw0000gn/T/go-build3659010178/b001/exe/assign_interface_hard myTextFile.txt]'

If I have this code: fmt.Println(os.Args[1])
	This will print my 1 argument which will be 'myTextFile.txt'

*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])
	os.Open(os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
