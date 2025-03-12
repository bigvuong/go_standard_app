package logs

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Log(message string) {
	logger.Println(message)
}
