package main

import (
	"bufio"
	"log"
	"os"
	"textcom/file"
	"textcom/handler"
)

func main() {
	f, err := file.New(1)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(f)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	h.Run(text)
}
