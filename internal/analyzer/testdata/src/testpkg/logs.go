package testpkg

import (
	"log"
	"log/slog"
)

func test() {
	log.Println("Hello world")        // want "log message should start with a lowercase letter"
	log.Println("Привет мир")          // want "log message should use only English \\(ASCII\\) characters"
	log.Println("Password: 1234")      // want "log message may contain sensitive data \\(e.g., 'password'\\)"
	log.Println("Hello!!!")            // want "log message should not contain special characters or punctuation"
	
	slog.Info("starting server")        // OK
	slog.Error("failed to connect")     // OK
	slog.Info("Hello world")            // want "log message should start with a lowercase letter"
}