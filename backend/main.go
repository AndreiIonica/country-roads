package main

import (
	"backend/roralis/config"
	"backend/roralis/dic"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Launches the app.
// Looks for configuration and errors out if not found.
// Starts up dependecy injection container.
// Starts the web server.
func main() {

	if err := config.Load(); err != nil {
		fmt.Printf("Failed to read config:\n     %v", err)
		os.Exit(1)
	}

	if _, err := dic.InitContainer(); err != nil {
		fmt.Printf("Failed to start Dependency Injection Conatiner:\n    %v", err)
		os.Exit(1)
	}

	app, err := config.StartApp()
	if err != nil {
		fmt.Printf("Failed to asemble Gin server:\n     %v", err)
		os.Exit(1)
	}
	// By default, Gin uses the PORT enviroment variables

	fmt.Printf("Server running on localhost:%v\n", viper.GetString("PORT"))
	err = app.Run(":" + viper.GetString("PORT"))
	if err != nil {
		fmt.Printf("Failed to start Gin server:\n    %v", err)
		os.Exit(1)
	}

}
