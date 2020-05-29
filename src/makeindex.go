package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("index.jsp")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	v := os.Getenv("name")
	s := `<html>
<body>
<h2>Hello `+ v +` the Server x2  is running!</h2>
<h1>Test to put the artifact on Bitbucket repository </h1>
</body>
</html>`
	fmt.Fprintf(file, s)
}
