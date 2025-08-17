// my own code
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// type logWriter struct{}

// func main() {
// 	lw := logWriter{}
// 	fileName := os.Args[1]
// 	openedFile, err := os.Open(fileName)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	io.Copy(lw, openedFile)
// }

// func (logWriter) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	return len(bs), nil
// }

// instructor code
// using os.Stdout instead of custom writer function
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
