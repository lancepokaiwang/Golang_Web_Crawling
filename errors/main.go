package errors

import (
	"log"
	"os"
	"time"
)

// init deals with setting up log information.
func init() {
	fileName := "logs/" + GetDateTime() + ".txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

// Fatal records and outputs the error message then shut the entire application down immediately.
func Fatal(message string) {
	log.Println(message)
	os.Exit(0)
}

// ContextLog records and outputs the given message.
func ContextLog(message string) {
	log.Println(message)
}

// GetDateTime returns the current datetime in string format.
func GetDateTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
