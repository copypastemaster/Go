package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// log.Fatal(http.ListenAndServe(":3001", http.HandlerFunc(MyGreetingHandler)))
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var writer io.Writer = file
	n, err := writer.Write([]byte("Hello, wrodl!"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Wrote %d bytes to file\n", n)

	exampleFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer exampleFile.Close()

	var reader io.Reader = exampleFile
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro reading file:", err)
			return
		}

		_, err = os.Stdout.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to stdout:", err)
			return
		}
	}
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
