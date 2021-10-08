package logging

import (
	"log"
	"os"
)

func LogToFile(msg string, level string) {
	file, err := os.OpenFile("logging/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	log.SetOutput(file)
	if level == "fatal" {
		log.Fatal(msg)
	} else {
		log.Println(msg)
	}
}
