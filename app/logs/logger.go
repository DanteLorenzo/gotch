package logs

import (
    "os"
    "github.com/fatih/color" // Import the color package for colored output
)

// LogInfo logs informational messages to the console with the specified color.
// It takes a message string and a color pointer to determine the color of the log output.
func LogInfo(message string, color *color.Color) {
    color.Printf("[INFO] %s\n", message) // Print the message with an [INFO] prefix in the specified color
}

// LogFatal logs error messages to the console in red and then terminates the program.
// It takes a format string and a variadic number of arguments to format the error message.
func LogFatal(format string, args ...interface{}) {
    color.New(color.FgRed).Printf("[ERROR] "+format+"\n", args...) // Print the error message in red
    os.Exit(1) // Exit the program with a status code of 1
}
